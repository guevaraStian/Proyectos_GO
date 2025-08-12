
// Se crea el ajax con la informacion que puede venir da la base de datos

const ctx = document.getElementById('myChart').getContext('2d');
const chart = new Chart(ctx, {
    type: 'bar',
    data: {
        labels: ['Usuarios', 'Visitantes'],
        datasets: [{
            label: 'Datos actuales',
            data: [111, 222],
            backgroundColor: ['#36a2eb', '#ff6384']
        }]
    },
    options: {
        responsive: true,
        plugins: {
            legend: {
                position: 'top'
            }
        }
    }
});