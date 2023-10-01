import json
from datetime import datetime
from pathlib import Path
import time

import openai

def get_fine_tuned_models(limit: int = 10):
    return openai.FineTuningJob.list(limit)['data']


def train_model(fp):
    return create_fine_tuning_job(training_file=fp)


def create_fine_tuning_job(training_file, model="gpt-3.5-turbo"):
    return openai.FineTuningJob.create(training_file=training_file, model=model)


def get_fine_tuning_jobs():
    return openai.FineTuningJob.list()["data"]

def get_fine_tuning_job(id):
    return next((item for item in get_fine_tuning_jobs() if item.get('id') == id), None)


def get_files():
    return openai.File.list()["data"]


def create_file(file_path, purpose="fine-tune"):
    return openai.File.create(
        file=open(file_path, "rb"),
        purpose=purpose
    )


def convert_to_jsonl(data):
    # Define the output file name
    current_time = datetime.now().strftime("%Y-%m-%d_%H-%M-%S")
    jsonl_filename = Path.cwd() / f"output_{current_time}.jsonl"
    # Write JSON data to JSONL file
    with open(jsonl_filename, "w") as jsonl_file:
        for entry in data:
            # Convert dictionary to JSON string and write to the file
            json_line = json.dumps(entry)
            jsonl_file.write(json_line + "\n")

    return jsonl_filename


def poll_for_job_completion(id) -> str:
    job = get_fine_tuning_job(id)
    if job:
        while not job["status"] == "completed":
            job = get_fine_tuning_job(id)
            time.sleep(30)

        return job["fine_tuned_model"]
    raise ValueError("Theres not job with the specified ID")

