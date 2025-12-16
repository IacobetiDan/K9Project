#!/usr/bin/env bash

# Directorul de bază al proiectului pe Desktop
BASE_DIR="$HOME/onedrive/Desktop/k8s-project"

echo "Crez structura Kubernetes în: $BASE_DIR"

# Foldere principale
mkdir -p "$BASE_DIR"/{docs,app,manifests,overlays,charts,scripts,ci-cd}

# Subfoldere pentru manifests
mkdir -p "$BASE_DIR"/manifests/{namespaces,deployments,services,ingresses,configmaps,secrets,storage,crds}

# Subfoldere pentru overlays (medii)
mkdir -p "$BASE_DIR"/overlays/{dev,test,prod}

# Fișiere de bază
touch "$BASE_DIR/README.md"
touch "$BASE_DIR/.gitignore"

echo "Structura a fost creată."
echo "Poți găsi proiectul în: $BASE_DIR"
