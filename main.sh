DEBIAN_FRONTEND=noninteractive

# Clone Upstream
mkdir -p ./pikman
cp -rvf ./* ./pikman || echo
cd ./pikman

# Get build deps
apt-get build-dep ./ -y

# Build package
dpkg-buildpackage

# Move the debs to output
cd ../
mkdir -p ./output
mv ./*.deb ./output/

