from fastapi import FastAPI
from fastapi.responses import JSONResponse
import math

app = FastAPI()

app.get("/ping")
async def ping():
    return JSONResponse(content={"message": "pong"})

if __name__ == "__main__":
    import uvicorn
    uvicorn.run("main:app", host="0.0.0.0", port=8079, workers=4)