document.addEventListener('DOMContentLoaded', () => {

    const searchForm = document.getElementById('searchForm');
    const orderUidInput = document.getElementById('orderUidInput');
    const jsonOutput = document.getElementById('jsonOutput');
    const errorMessage = document.getElementById('errorMessage');

    searchForm.addEventListener('submit', async (event) => {
        event.preventDefault();

        const orderUid = orderUidInput.value.trim();
        if (!orderUid) {
            return;
        }

        // Прячем предыдущие результаты перед новым запросом
        jsonOutput.style.display = 'none';
        errorMessage.style.display = 'none';

        try {
            const response = await fetch(`/order/${orderUid}`);

            if (response.ok) {
                const data = await response.json();

                jsonOutput.textContent = JSON.stringify(data, null, 2);
                jsonOutput.style.display = 'block';

            } else if (response.status === 404) {
                errorMessage.textContent = 'Заказ не найден.';
                errorMessage.style.display = 'block';
            } else {
                errorMessage.textContent = `Ошибка сервера: ${response.status} ${response.statusText}`;
                errorMessage.style.display = 'block';
            }

        } catch (error) {
            console.error('Ошибка при выполнении запроса:', error);
            errorMessage.textContent = 'Не удалось подключиться к серверу. Проверьте сеть.';
            errorMessage.style.display = 'block';
        }
    });
});