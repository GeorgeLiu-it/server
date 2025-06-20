# Environment setup

## System setup

### Ubuntu image

[Download](https://ubuntu.com/download/server)

### Ubuntu installation

I am using VMWare to install ubuntu.

![image-20250618153228165](C:\Users\liushanshan\AppData\Roaming\Typora\typora-user-images\image-20250618153228165.png)

**My info**

name: george

server name: blog_server

username: george

**Update ubuntu packages**

```bash
sudo apt-get update
```

## CA trust

### Zscaler

```bash
sudo cp [mycert.crt] /usr/local/share/ca-certificates/
sudo update-ca-certificates
```

## Software setup

### VIM

```bash
sudo apt-get install vim
```

### Docker with docker-compose

```bash
# remove the conflicting packages
for pkg in docker.io docker-doc docker-compose docker-compose-v2 podman-docker containerd runc; do sudo apt-get remove $pkg; done

# Add Docker's official GPG key:
sudo apt-get update
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

# Add the repository to Apt sources:
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "${UBUNTU_CODENAME:-$VERSION_CODENAME}") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update

# install docker packages
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

### Alternative solution

**Download the docker installation package and install separately**

[Package link](https://download.docker.com/linux/ubuntu/dists/noble/pool/stable/amd64/)

```bash
dpkg -i ./*.deb
```

![image-20250619083809799](C:\Users\liushanshan\AppData\Roaming\Typora\typora-user-images\image-20250619083809799.png)

## Infrastructure

### Blog backend

system: ubuntu 24.04.2

ip: 192.168.68.231

### Rag backend

system: ubuntu 24.04.2

ip: 192.168.68.232

### Blog frontend

system: ubuntu 24.04.2

ip: 192.168.68.233

## Reference

### Change ip

```bash
# Check the NIC
ip link show

# Modify the ip section in the netplan configuration yaml
sudo vim /etc/netplan/50-cloud-init.yaml

# Apply the changes
sudo netplan apply
```

### Change hostname

```bash
# Check current hostname
hostnamectl

# Change hostname
sudo hostnamectl set-hostname [hostname]

# Update hosts
sudo vim /etc/hosts
# add 127.0.1.1 [hostname]
```

