fetch('https://raw.githubusercontent.com/riviox/GitMan/main/packages.json')
    .then(response => response.json())
    .then(data => {
    const packagesList = document.getElementById('packages-list');
    data.forEach(package => {
        const packageItem = document.createElement('div');
        const packageName = document.createElement('a');
        packageName.textContent = package.name;
        packageName.href = package.repository;
        packageName.target = "_blank";
        packageItem.appendChild(packageName);
        packagesList.appendChild(packageItem);
    });
    })
.catch(error => console.error('Błąd podczas pobierania danych:', error));