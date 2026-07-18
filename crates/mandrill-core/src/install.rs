use std::path::PathBuf;
use anyhow::Result;
use flate2::read::GzDecoder;
use std::fs::File;
use tar::Archive;

pub fn get_mandrill_dir() -> PathBuf {
    dirs::home_dir()
        .expect("could not find home directory")
        .join(".mandrill")
}

pub fn get_bin_dir() -> PathBuf {
    get_mandrill_dir().join("bin")
}

pub fn get_packages_dir() -> PathBuf {
    get_mandrill_dir().join("packages")
}

pub fn get_registry_dir() -> PathBuf {
    get_mandrill_dir().join("registry")
}

pub fn get_lock_path() -> PathBuf {
    get_mandrill_dir().join("lock.json")
}

pub fn ensure_dirs() -> Result<()> {
    let dirs = vec![get_bin_dir(), get_packages_dir(), get_registry_dir()];
    for dir in dirs {
        std::fs::create_dir_all(dir)?;
    }
    Ok(())
}

pub fn extract_tar_gz(archive_path: &PathBuf, dest: &PathBuf) -> Result<()> {
    let file = File::open(archive_path)?;
    let dec = GzDecoder::new(file);
    let mut archive = Archive::new(dec);
    archive.unpack(dest)?;
    Ok(())
}

pub fn symlink_bin(src: &PathBuf, bin_dir: &PathBuf) -> Result<()> {
    let filename = src.file_name().expect("no filename");
    let link_path = bin_dir.join(filename);
    if link_path.exists() {
        std::fs::remove_file(&link_path)?;
    }
    std::os::unix::fs::symlink(src, &link_path)?;
    Ok(())
}
