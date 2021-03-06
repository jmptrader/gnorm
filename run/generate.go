package run // import "gnorm.org/gnorm/run"

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/pkg/errors"

	"gnorm.org/gnorm/environ"
	"gnorm.org/gnorm/run/data"
)

// Generate reads your database, gets the schema for it, and then generates
// files based on your templates and your configuration.
func Generate(env environ.Values, cfg *Config) error {
	info, err := cfg.Driver.Parse(env.Log, cfg.ConnStr, cfg.Schemas, makeFilter(cfg.IncludeTables, cfg.ExcludeTables))
	if err != nil {
		return err
	}
	db, err := makeData(env.Log, info, cfg)
	if err != nil {
		return err
	}
	if len(cfg.SchemaPaths) == 0 {
		env.Log.Println("No SchemaPaths specified, skipping schemas.")
	} else {
		if err := generateSchemas(env, cfg, db); err != nil {
			return err
		}
	}
	if len(cfg.EnumPaths) == 0 {
		env.Log.Println("No EnumPath specified, skipping enums.")
	} else {
		if err := generateEnums(env, cfg, db); err != nil {
			return err
		}
	}
	if len(cfg.TablePaths) == 0 {
		env.Log.Println("No table path specified, skipping tables.")
	} else {
		if err := generateTables(env, cfg, db); err != nil {
			return err
		}
	}
	return nil
}

func generateSchemas(env environ.Values, cfg *Config, db *data.DBData) error {
	for _, schema := range db.Schemas {
		fileData := struct{ Schema string }{Schema: schema.Name}
		contents := data.SchemaData{
			Schema: schema,
			DB:     db,
			Config: cfg.ConfigData,
			Params: cfg.Params,
		}
		for _, target := range cfg.SchemaPaths {
			env.Log.Printf("Generating output for schema %v", schema.Name)
			if err := genFile(env, fileData, contents, target, cfg.PostRun); err != nil {
				return errors.WithMessage(err, "generating file for schema "+schema.Name)
			}
		}
	}
	return nil
}

func generateEnums(env environ.Values, cfg *Config, db *data.DBData) error {
	for _, schema := range db.Schemas {
		for _, enum := range schema.Enums {
			fileData := struct{ Schema, Enum string }{Schema: schema.Name, Enum: enum.Name}
			contents := data.EnumData{
				Enum:   enum,
				DB:     db,
				Config: cfg.ConfigData,
				Params: cfg.Params,
			}
			for _, target := range cfg.EnumPaths {
				if err := genFile(env, fileData, contents, target, cfg.PostRun); err != nil {
					env.Log.Printf("Generating output for enum %v", enum.Name)
					return errors.WithMessage(err, "generating file for enum "+enum.Name)
				}
			}
		}
	}
	return nil
}

func generateTables(env environ.Values, cfg *Config, db *data.DBData) error {
	for _, schema := range db.Schemas {
		for _, table := range schema.Tables {
			contents := data.TableData{
				Table:  table,
				DB:     db,
				Config: cfg.ConfigData,
				Params: cfg.Params,
			}
			fileData := struct{ Schema, Table string }{Schema: schema.Name, Table: table.Name}
			for _, target := range cfg.TablePaths {
				if err := genFile(env, fileData, contents, target, cfg.PostRun); err != nil {
					env.Log.Printf("Generating output for table %v", table.Name)
					return errors.WithMessage(err, "generating file for table "+table.Name)
				}
			}
		}
	}
	return nil
}

func genFile(env environ.Values, filedata, contents interface{}, target OutputTarget, postrun []string) error {
	buf := &bytes.Buffer{}
	err := target.Filename.Execute(buf, filedata)
	if err != nil {
		return errors.WithMessage(err, "failed to run Filename template")
	}
	outputPath := buf.String()

	if err := os.MkdirAll(filepath.Dir(outputPath), 0700); err != nil {
		return errors.WithMessage(err, "error creating template output directory")
	}

	outbuf := &bytes.Buffer{}
	if err := target.Contents.Execute(outbuf, contents); err != nil {
		return errors.WithMessage(err, "failed to run contents template")
	}
	if err := ioutil.WriteFile(outputPath, outbuf.Bytes(), 0600); err != nil {
		return errors.Wrapf(err, "error writing generated file %q", outputPath)
	}
	if len(postrun) > 0 {
		return doPostRun(env, outputPath, postrun)
	}
	return nil
}

func doPostRun(env environ.Values, file string, postrun []string) error {
	newenv := make(map[string]string, len(env.Env)+1)
	for k := range env.Env {
		newenv[k] = env.Env[k]
	}
	run := make([]string, len(postrun))
	newenv["GNORMFILE"] = file
	conv := func(s string) string { return newenv[s] }
	for x, s := range postrun {
		run[x] = os.Expand(s, conv)
	}
	var cmd *exec.Cmd
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if len(run) > 1 {
		cmd = exec.CommandContext(ctx, run[0], run[1:]...)
	} else {
		cmd = exec.CommandContext(ctx, run[0])
	}
	cmd.Stderr = env.Stderr
	cmd.Stdout = env.Stdout
	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "error running postrun command %q", run)
	}
	return nil
}
