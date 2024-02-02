<h1 align="center" id="title">PDF Compressor</h1>

<p align="center"><img src="https://socialify.git.ci/suhail34/PDF-Compressor/image?language=1&amp;owner=1&amp;name=1&amp;stargazers=1&amp;theme=Light" alt="project-image"></p>

<p id="description">The PDF Compressor project is a system for compressing PDF files and managing them using a combination of MongoDB Kafka and a Golang-based web service. This README provides an overview of the project its components and how to set it up.</p>

<h2>💻 Components</h2>

<h3>1. Web Service (Golang)</h3>

<ul>
    <li><strong>Description:</strong> Handles file uploads, compression, and file downloads.</li>
    <li><strong>Dependencies:</strong> Gin (HTTP framework), MongoDB Go driver, Confluent Kafka Go library.</li>
</ul>

<h3>2. Kafka (Message Broker)</h3>

<ul>
    <li><strong>Description:</strong> Manages asynchronous communication between the web service and the consumer service.</li>
    <li><strong>Dependencies:</strong> Confluent Kafka.</li>
</ul>

<h3>3. MongoDB (Database)</h3>

<ul>
    <li><strong>Description:</strong> Stores compressed PDF files and associated metadata.</li>
    <li><strong>Dependencies:</strong> MongoDB, GridFS.</li>
</ul>

<h3>4. Consumer Service (Python)</h3>

<ul>
    <li><strong>Description:</strong> Consumes messages from Kafka, indicating successful compression.</li>
    <li><strong>Dependencies:</strong> Confluent Kafka.</li>
</ul>

<h3>5. Database Cleaner Service (Python)</h3>

<ul>
    <li><strong>Description:</strong> Scheduled job for deleting old files from MongoDB GridFS.</li>
    <li><strong>Dependencies:</strong> MongoDB, GridFS</li>
</ul>

<br>

  
<h2>🧐 Features</h2>

Here're some of the project's best features:

*   Upload PDF files for compression.
*   Compress PDF files using PyMuPDF
*   Store compressed PDF files in MongoDB using GridFS.
*   Asynchronous processing of compression tasks with Kafka.
*   Download compressed PDF files.

<h2>🛠️ Installation Steps:</h2>

<p>1. Fork The Repository</p>

<p>2. Clone the repository:</p>

```
git clone https://github.com/your-username/pdf-compressor.git
```

<p>3. Navigate to the project directory:</p>

```
cd pdf-compressor
```

<p>4. Add the helm bitnami repo:</p>

```
helm repo add bitnami https://charts.bitnami.com/bitnami
```

<p>5. Install the Confluentic Kafka service using below command:</p>

```
helm upgrade --install kafka-release bitnami/kafka --set persistence.size=8GilogPersistence.size=8GivolumePermissions.enabled=truepersistence.enabled=truelogPersistence.enabled=trueserviceAccount.create=truerbac.create=true --version 23.0.7 -f Helm_charts/Kafka/values.yaml
```

<p>6. Install MongoDB service using command:</p>

```
helm install mongo Helm_charts/MongoDB -f Helm_chart/MongoDB/values.yaml
```

<p>7. Run frontend service using command</p>

```
kubectl apply -f frontend-service/manifests/
```

<p>8. Run producer service using command</p>

```
kubectl apply -f producer-service/manifests/
```

<p>9. Run compressor service using command</p>

```
kubectl apply -f compressor-service/manifests/
```

<p>10. Run DB Cleaner Service using command</p>

```
kubectl apply -f DBClean-service/manifests/
```
  
<h2>💻 Built with</h2>

Technologies used in the project:

*   Golang
*   Kafka
*   Helm Charts
*   MongoDB
*   Python
*   Html
*   Bootstrap
*   Javascript

<h2>🛡️ License:</h2>

This project is licensed under the MIT
