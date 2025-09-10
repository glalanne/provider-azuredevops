#!/usr/bin/env python3

import os
import sys
import glob
from pathlib import Path


if __name__ == "__main__":
    directory=".work/microsoft/azuredevops/website/docs/r"
    files=glob.glob(f"{directory}/*.markdown")

    resources=[]
    for file_path in files:
        base_name=Path(Path(file_path).stem).stem
  
        resources.append(base_name)

        os.makedirs(f"config/namespaced/{base_name}", exist_ok=True)
        if not os.path.exists(f"config/namespaced/{base_name}/config.go"):
            with open(f"config/namespaced/{base_name}/config.go", "w") as f:

                f.write(f'package {base_name}\n\n')
                f.write('import "github.com/crossplane/upjet/v2/pkg/config"\n\n')
                f.write('// Configure configures individual resources by adding custom ResourceConfigurators.\n')
                f.write('func Configure(p *config.Provider) {\n')
                f.write(f'    p.AddResourceConfigurator("azuredevops_{base_name}", func(r *config.Resource) {{\n')
                f.write('         r.ShortGroup = ""\n')
                f.write('    })\n')
                f.write('}\n')


            print(f"{base_name} created.")

        os.makedirs(f"config/cluster/{base_name}", exist_ok=True)
        if not os.path.exists(f"config/cluster/{base_name}/config.go"):
            with open(f"config/cluster/{base_name}/config.go", "w") as f:

                f.write(f'package {base_name}\n\n')
                f.write('import "github.com/crossplane/upjet/v2/pkg/config"\n\n')
                f.write('// Configure configures individual resources by adding custom ResourceConfigurators.\n')
                f.write('func Configure(p *config.Provider) {\n')
                f.write(f'    p.AddResourceConfigurator("azuredevops_{base_name}", func(r *config.Resource) {{\n')
                f.write('         r.ShortGroup = ""\n')
                f.write('    })\n')
                f.write('}\n')


            print(f"{base_name} created.")


    print(f"Total resources: {len(resources)}")
    print("------")

    for resource in resources:
        print(f"\"github.com/glalanne/provider-azuredevops/config/cluster/{resource}\"")
   

    print("------")
    for resource in resources:
        print(f"\tProviderConfiguration.AddConfig({resource}.Configure)")

    print("------")

    for resource in resources:
        print(f"\"github.com/glalanne/provider-azuredevops/config/namespaced/{resource}\"")
   

    print("------")
    for resource in resources:
        print(f"\tProviderConfiguration.AddConfig({resource}.Configure)")

    print("------")
    for resource in resources:
        print(f"\t\"azuredevops_{resource}\":                                        config.IdentifierFromProvider,")

    print("------")