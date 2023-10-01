import os

from fastapi import FastAPI
import openai

from . import utils

app = FastAPI()

# Set your OpenAI API key here
OPENAI_API_KEY = os.getenv("OPENAI_API_KEY")
openai.api_key = OPENAI_API_KEY


MODELS = {}

@app.post("/train/{org_id}/")
async def train_model(org_id: str,data: dict):
    file_name = utils.convert_to_jsonl(data)
    file_id = utils.create_file(file_name)["id"]
    utils.create_fine_tuning_job(file_id)
    ftm =utils.poll_for_job_completion(file_id)

    # should persist this in a database instead.
    MODELS[org_id] = ftm

    return {"message": f"Model successfully trained"}

