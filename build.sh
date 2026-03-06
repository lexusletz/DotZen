#!/bin/bash

BINARY_NAME="dotzen"
BIN_DIR="bin"
DIST_DIR="dist"
VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "dev")

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

show_help() {
    echo -e "${GREEN}DotZen Build Script${NC}"
    echo ""
    echo "Usage: ./build.sh [options]"
    echo ""
    echo "Options:"
    echo "  -a, --all       Compile for all the platforms"
    echo "  -c, --clean     Cleans the generated files"
    echo "  -r, --release   Create release with compressed files"
    echo "  -i, --install   Install the binary locally"
    echo "  -h, --help      Shows this help page"
}

build_current() {
    echo -e "${YELLOW}Compiling for the current platform...${NC}"

    mkdir -p "$BIN_DIR"

    go build -ldflags "-X main.Version=$VERSION" -o "$BIN_DIR/$BINARY_NAME" ./cmd/dotzen

    if [ $? -eq 0 ]; then
        echo -e "${GREEN}Compilation success: $BIN_DIR/$BINARY_NAME${NC}"
    else
        echo -e "${RED}Error in compilation${NC}"
        exit 1
    fi
}

build_all() {
    echo -e "${YELLOW}Compiling for all platforms...${NC}"

    mkdir -p "$BIN_DIR"

    platforms=("darwin/amd64" "darwin/arm64" "linux/amd64" "linux/arm64" "windows/amd64")

    for platform in "${platforms[@]}"; do
        os=$(echo "$platform" | cut -d'/' -f1)
        arch=$(echo "$platform" | cut -d'/' -f2)
        output_name="$BINARY_NAME-$os-$arch"

        if [ "$os" = "windows" ]; then
            output_name="$output_name.exe"
        fi

        echo -e "${CYAN}Compiling for $os/$arch...${NC}"

        GOOS=$os GOARCH=$arch go build -ldflags "-X main.Version=$VERSION" -o "$BIN_DIR/$output_name" ./cmd/dotzen

        if [ $? -ne 0 ]; then
            echo -e "${RED}Error compiling for $os/$arch${NC}"
            exit 1
        fi
    done

    echo -e "${GREEN}All compilations completed${NC}"
}

create_release() {
    build_all

    echo -e "${YELLOW}Creating release...${NC}"
    mkdir -p "$DIST_DIR"

    platforms=("darwin/amd64" "darwin/arm64" "linux/amd64" "linux/arm64" "windows/amd64")

    for platform in "${platforms[@]}"; do
        os=$(echo "$platform" | cut -d'/' -f1)
        arch=$(echo "$platform" | cut -d'/' -f2)
        binary_name="$BINARY_NAME-$os-$arch"

        if [ "$os" = "windows" ]; then
            binary_name="$binary_name.exe"
        fi 

        archive_name="$BINARY_NAME-$VERSION-$os-$arch"

        if [ "$os" = "windows" ]; then
            zip -j "$DIST_DIR/$archive_name.zip" "$BIN_DIR/$binary_name" README.md
        else
            tar -czf "$DIST_DIR/$archive_name.tar.gz" -C "$BIN_DIR" "$binary_name" -C ../. README.md
        fi

        echo -e "${CYAN}Created: $DIST_DIR/$archive_name${NC}"
    done

    echo -e "${GREEN}Release completed in $DIST_DIR/${NC}"
}

install_local() {
    build_current

    echo -e "${YELLOW}Installing $BINARY_NAME...${NC}"
    sudo cp "$BIN_DIR/$BINARY_NAME" /usr/local/bin/
    sudo -r "${GREEN}$BINARY_NAME installed in /usr/local/bin/${NC}"
}

clean_files() {
    echo -e "${YELLOW}Cleaning generated files...${NC}"
    rm -rf "$BIN_DIR" "$DIST_DIR"
    echo -e "${GREEN}Cleaning completed${NC}"
}

chmod +x "$0"

case "${1:-}" in
    -a|--all)
        build_all
        ;;
    -c|--clean)
        clean_files
        ;;
    -r|--release)
        create_release
        ;;
    -i|--install)
        install_local
        ;;
    -h|--help)
        show_help
        ;;
    "")
        build_current
        ;;
    *)
        echo -e "${RED}Unknown option: $1${NC}"
        show_help
        exit 1
        ;;
esac
