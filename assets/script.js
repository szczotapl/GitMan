fetch('https://raw.githubusercontent.com/riviox/GitMan/main/packages.json')
    .then(response => response.json())
    .then(data => {
        const packagesList = document.getElementById('packages-list');
        let column;
        data.forEach((package, index) => {
            if (index % 5 === 0) {
                column = document.createElement('div');
                packagesList.appendChild(column);
            }

            const packageItem = document.createElement('div');
            const packageName = document.createElement('a');
            packageName.textContent = package.name;
            packageName.href = package.repository;
            packageName.target = "_blank";
            packageItem.appendChild(packageName);
            column.appendChild(packageItem);
        });
    })
    .catch(error => console.error('Error:', error));
