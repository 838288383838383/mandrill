use serde::Deserialize;
use std::path::Path;
use mandrill_core::manifest::PackageManifest;
use anyhow::Result;

#[derive(Debug, Deserialize)]
pub struct RegistryIndex {
    pub packages: std::collections::HashMap<String, PackageEntry>,
}

#[derive(Debug, Deserialize)]
pub struct PackageEntry {
    pub version: String,
    pub manifest: String,
}

impl RegistryIndex {
    pub fn from_file(path: &Path) -> Result<Self> {
        let content = std::fs::read_to_string(path)?;
        let index: RegistryIndex = toml::from_str(&content)?;
        Ok(index)
    }

    pub fn search(&self, query: &str) -> Vec<(&String, &PackageEntry)> {
        self.packages
            .iter()
            .filter(|(name, _)| name.contains(query))
            .collect()
    }

    pub fn get_manifest(&self, registry_path: &Path, name: &str) -> Result<PackageManifest> {
        let entry = self.packages.get(name)
            .ok_or_else(|| anyhow::anyhow!("package '{}' not found in registry", name))?;
        let manifest_path = registry_path.join(&entry.manifest);
        PackageManifest::from_file(&manifest_path)
    }
}
