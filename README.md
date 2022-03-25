# Go Study

### Test On Save
- vscode 의 runonsave 플러그인 이용
- .vscode 폴더를 프로젝트 최상단에 두고, settings.json 에 다음을 추가한다.

```json
{
    "emeraldwalk.runonsave": {
        "autoClearConsole": true,
        "commands": [
            {
                "match": ".go",
                "cmd": "make run"
            }
        ]
    }
}
```

- 그 후 하단 output 창에 결과를 확인할 수 있다.