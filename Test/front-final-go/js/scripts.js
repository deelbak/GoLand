function showNotification() {
    // Проверяем, поддерживает ли браузер уведомления
    if (!("Notification" in window)) {
      alert("Этот браузер не поддерживает уведомления");
    } else if (Notification.permission === "granted") {
      // Если разрешено показывать уведомления, выводим сообщение
      new Notification("Уведомление", {
        body: "Событие произошло",
        icon: "https://example.com/icon.png",
      });
    } else if (Notification.permission !== "denied") {
      // Если еще не получили разрешение, запрашиваем его
      Notification.requestPermission().then(function (permission) {
        if (permission === "granted") {
          // Если разрешено показывать уведомления, выводим сообщение
          new Notification("Уведомление", {
            body: "Событие произошло",
            icon: "https://example.com/icon.png",
          });
        }
      });
    }
  }
// Получаем ссылки на элементы кнопки и модального окна
var openModalBtn = document.getElementById("open-modal-btn");
var successModal = document.getElementById("success-modal");

// Получаем ссылку на кнопку закрытия модального окна
var closeModalBtn = successModal.getElementsByClassName("close")[0];

// Функция, вызываемая при нажатии на кнопку "Зарегистрироваться"
function openModal() {
  successModal.style.display = "block"; // отображаем модальное окно
}

// Функция, вызываемая при нажатии на кнопку закрытия модального окна
function closeModal() {
  successModal.style.display = "none"; // скрываем модальное окно
}

// Назначаем обработчики событий для кнопок
openModalBtn.onclick = openModal;
closeModalBtn.onclick = closeModal;
// получаем элемент с сообщением
const successMessage = document.getElementById('success-message');

// функция, которая показывает сообщение
function showSuccessMessage() {
  successMessage.style.display = 'block';
}

// функция, которая скрывает сообщение
function hideSuccessMessage() {
  successMessage.style.display = 'none';
}
  