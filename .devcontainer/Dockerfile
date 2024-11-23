FROM --platform=$BUILDPLATFORM golang:1.23.1-alpine 
LABEL maintainer="@k33g_org"

ARG USER_NAME=${USER_NAME}

# Install required packages
RUN apk update && apk add --no-cache \
    curl \
    wget \
    sudo \
    sshpass \
    jq \
    bash \
    git \
    shadow # Needed for usermod command

# Create user and add to sudo group
RUN adduser -D ${USER_NAME} && \
    addgroup sudo && \
    adduser ${USER_NAME} sudo

# Configure sudo without password for sudo group
RUN echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

# Set the working directory
WORKDIR /home/${USER_NAME}

# Install Oh My Bash
RUN bash -c "$(curl -fsSL https://raw.githubusercontent.com/ohmybash/oh-my-bash/master/tools/install.sh)" --unattended

# Copy Oh My Bash configuration to user's home
RUN cp /root/.bashrc /home/${USER_NAME}/.bashrc && \
    cp -r /root/.oh-my-bash /home/${USER_NAME}/.oh-my-bash

# Set the user as the owner of the working directory and Oh My Bash files
RUN chown -R ${USER_NAME}:${USER_NAME} /home/${USER_NAME}

# Switch to the regular user
USER ${USER_NAME}

# Set default shell to bash
SHELL ["/bin/bash", "-c"]