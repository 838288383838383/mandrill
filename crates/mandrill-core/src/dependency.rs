use std::collections::HashMap;
use crate::manifest::PackageManifest;

#[derive(Debug)]
pub struct DependencyGraph {
    pub packages: HashMap<String, PackageManifest>,
}

impl DependencyGraph {
    pub fn new() -> Self {
        Self {
            packages: HashMap::new(),
        }
    }

    pub fn add(&mut self, manifest: PackageManifest) {
        self.packages
            .insert(manifest.package.name.clone(), manifest);
    }

    pub fn resolve_order(&self) -> Vec<String> {
        let mut resolved = Vec::new();
        let mut visited = std::collections::HashSet::new();

        for name in self.packages.keys() {
            self.visit(name, &mut resolved, &mut visited);
        }
        resolved
    }

    fn visit(&self, name: &str, resolved: &mut Vec<String>, visited: &mut std::collections::HashSet<String>) {
        if visited.contains(name) {
            return;
        }
        visited.insert(name.to_string());

        if let Some(manifest) = self.packages.get(name) {
            if let Some(deps) = &manifest.dependencies {
                for dep in deps {
                    self.visit(dep, resolved, visited);
                }
            }
        }
        resolved.push(name.to_string());
    }
}
