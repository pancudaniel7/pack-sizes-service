function submitPackSizes() {
    const packSizeInputs = document.querySelectorAll('.form-group:nth-of-type(1) input[type="number"]');
    let packSizes = Array.from(packSizeInputs).map(input => parseInt(input.value, 10)).filter(value => value > 0);

    fetch('/set-pack-sizes', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({sizes: packSizes})
    })
        .then(response => response.json())
        .then(data => console.log('Success:', data))
        .catch(error => console.error('Error:', error));
}

function calculatePacks() {
    const orderQty = document.getElementById('items').value;

    fetch(`/calculate-packs?orderQty=${orderQty}`, {
        method: 'GET'
    })
        .then(response => response.json())
        .then(data => {
            console.log('Success:', data);
            // Clear previous pack display
            const packDisplay = document.querySelector('.form-group:last-of-type');
            packDisplay.innerHTML = `
                <div class="row">
                    <div><label>Pack</label></div>
                    <div><label>Quantity</label></div>
                </div>
            `;

            data.forEach(pack => {
                const packRow = document.createElement('div');
                packRow.classList.add('row');
                packRow.innerHTML = `
                    <div><input type="text" readonly value="${pack.quantity}"></div>
                    <div><input type="text" readonly value="${pack.size}"></div>

                `;
                packDisplay.appendChild(packRow);
            });
        })
        .catch(error => console.error('Error:', error));
}

document.getElementById('submitPackSizes').addEventListener('click', submitPackSizes);
document.getElementById('calculatePacks').addEventListener('click', calculatePacks);