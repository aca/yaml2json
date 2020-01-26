YAML to JSON(y2j), JSON to YAML(j2y)
---

Two extremely simple wrapper of [ghodss/yaml](https://github.com/ghodss/yaml) to handle yaml. Zero config, get inputs only from stdin.  
For patching, querying YAML, Let the great 'jq' do all the hard work.

```
cat pod.yaml | y2j
cat pod.json | j2y

kubectl get pods -A -o json | j2y | y2j | jq .
```

#### Install
Using go mod, requires go>=1.13
```
git clone git@github.com:aca/yaml2json.git && cd yaml2json 
go install ./j2y
go install ./y2j
```



