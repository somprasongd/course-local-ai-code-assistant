# Open WebUI

## Installation

Docker: [doc](https://docs.openwebui.com/)

```bash
docker run -d -p 3000:8080 --add-host=host.docker.internal:host-gateway -v open-webui:/app/backend/data --name open-webui --restart always ghcr.io/open-webui/open-webui:main
```

## Demo

- [Create CRUD product API](/2.open-webui/demo/crud_api.md)
- [Natural language to SQL](/2.open-webui/demo/write_sql.md)
