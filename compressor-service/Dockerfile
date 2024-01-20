FROM python:3.10.13

WORKDIR /app

COPY . .

RUN pip install -r requirements.txt

Expose 8000

ENTRYPOINT ["python", "main.py"]
