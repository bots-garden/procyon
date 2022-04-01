# 🚧 work in progress
FROM gitpod/workspace-full

ARG GO_VERSION="1.17.5"
ARG TINYGO_VERSION="0.21.0"
#ARG RUST_VERSION="1.58.1"
ARG NODEJS_VERSION="17.7.2"
#ARG WABT_VERSION="1.0.28"
ARG WABT_VERSION="1.0.24"

USER gitpod

RUN sudo apt install libncurses5 libxkbcommon0 libtinfo5 libnss3-tools -y

# GoLang
RUN curl -sL https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer| bash

RUN ["/bin/bash", "-c", ". /home/gitpod/.gvm/scripts/gvm && gvm install go${GO_VERSION} -B"]
RUN ["/bin/bash", "-c", ". /home/gitpod/.gvm/scripts/gvm && gvm use go${GO_VERSION}"]

# TinyGo
RUN wget https://github.com/tinygo-org/tinygo/releases/download/v${TINYGO_VERSION}/tinygo_${TINYGO_VERSION}_amd64.deb
RUN sudo dpkg -i tinygo_${TINYGO_VERSION}_amd64.deb
RUN rm tinygo_${TINYGO_VERSION}_amd64.deb

RUN curl https://rustwasm.github.io/wasm-pack/installer/init.sh -sSf | sh 

# RustLang
# RUN bash -c "rustup toolchain install ${RUST_VERSION}"
RUN curl https://rustwasm.github.io/wasm-pack/installer/init.sh -sSf | sh 

# Nodejs
RUN bash -c 'VERSION="${NODEJS_VERSION}" && source $HOME/.nvm/nvm.sh && nvm install $VERSION && nvm use $VERSION &&nvm alias default $VERSION'
RUN echo "nvm use default &>/dev/null" >> ~/.bashrc.d/51-nvm-fix

RUN brew install httpie && \
    brew install hey && \
    brew install bat && \
    brew install exa && \
    brew install llvm && \
    brew install mkcert

RUN wget https://github.com/WebAssembly/wabt/releases/download/${WABT_VERSION}/wabt-${WABT_VERSION}-ubuntu.tar.gz && \
    tar -xf wabt-${WABT_VERSION}-ubuntu.tar.gz && \
    cd ./wabt-${WABT_VERSION}/bin && \
    ls && \
    sudo cp * /usr/local/bin 

RUN curl -sSf https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash && \
    curl https://get.wasmer.io -sSfL | sh && \
    curl https://wasmtime.dev/install.sh -sSf | bash
    
RUN rustup target add wasm32-wasi

# ------------------------------------
# Install the Suborbital CLI
# ------------------------------------
RUN brew tap suborbital/subo && \
    brew install subo
