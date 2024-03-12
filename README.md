# roadmap
Based on roadmap.sh, I follow my path with this repository 

## Build documentation

Documentation is build with antora. You can build it with docker: 
```shell
docker run -u $(id -u) -v $PWD:/antora:Z --rm -t antora/antora antora-playbook.yml
```
