var form = document.getElementById('fileUpload');

form.addEventListener("submit", (e) => {
  e.preventDefault();
  var fileInput = document.querySelector('input[name="files"]');
  var formData = new FormData();

  // Append each file to the FormData object
  for (var i = 0; i < fileInput.files.length; i++) {
    formData.append('files', fileInput.files[i]);
  }
  fetch('http://192.168.49.2:30001/api/files', {
    method: 'POST',
    body: formData,
  })
  .then(resp => resp.json())
  .then(data => {
    console.log('Success: ', data);
    generateDownloadBtns(data['resp'])
    generateDownloadAllBtn(data['resp'])
  })
  .catch((error) => {
    console.error('Error: ', error);
  })
  console.log(files);
})

function generateDownloadAllBtn(fileIdsMap) {
  var downloadAllBtnContainer = document.getElementById("downloadAllBtn");
  downloadAllBtnContainer.innerHTML = '';
  var downloadAllBtn = document.createElement('button');
  downloadAllBtn.innerText = 'Download All files';
  downloadAllBtn.className = 'btn btn-primary'
  downloadAllBtn.addEventListener('click', () => {
    for (var fileId in fileIdsMap) {
      downloadFile(fileId, fileIdsMap[fileId]);
    }
  });
  downloadAllBtnContainer.appendChild(downloadAllBtn)
}

function generateDownloadBtns(fileIdsMap) {
  var downloadButtonsContainer = document.getElementById("downloadButtons");
  downloadButtonsContainer.innerHTML = '';
  for (var fileId in fileIdsMap) {
    if (fileIdsMap.hasOwnProperty(fileId)) {
      var fileName = fileIdsMap[fileId]
      var downloadBtn = document.createElement('button');
      downloadBtn.innerText = 'Download File ' + fileName;
      downloadBtn.className = 'm-2 btn btn-secondary'
      downloadBtn.addEventListener('click', () => {
        downloadFile(fileId, fileIdsMap[fileId])
      });
      downloadButtonsContainer.appendChild(downloadBtn)
    }
  };
}

function downloadFile(fileId, fileName) {
  fetch(`http://192.168.49.2:30001/api/${fileId}`, {
    method: 'GET'
  })
  .then(resp => resp.blob())
  .then(blob => {
    var link = document.createElement('a');
    link.href = URL.createObjectURL(blob);
    link.download = fileName
    link.target = '_blank'
    link.click()
  })
  .catch(error => {
    console.error('Error Downloading file : ', error)
  });
}
