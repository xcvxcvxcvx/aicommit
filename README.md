## aicommit

Copy `aicommit` to the root of your project.  
Create a `prompt.txt` file with your prompt content.  
Add `aicommit` and `prompt.txt` to the local Git ignore:  

Add `aicommit` and `prompt.txt` to local git ignore

`$ echo aicommit >> .git/info/exclude`  
`$ echo prompt.txt >> .git/info/exclude`

### GIT hooks
Create or edit the .git/hooks/prepare-commit-msg hook:

``cd .git/hooks``  
``touch prepare-commit-msg``  
``chmod +x prepare-commit-msg``  

Add the following content:
```
#!/bin/sh
exec "$(git rev-parse --show-toplevel)/aicommit" "$@"
```

Add your gemini key:
`export GEMINI_API_KEY="{key}"`