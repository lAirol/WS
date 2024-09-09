let chatMain = document.getElementById('chat_main');
let messageInput = document.getElementById('message_input');
let sendButton = document.getElementById('send_button');

// Функция для добавления нового сообщения
function addMessage(message, isUser) {
    const newMessage = document.createElement('div');
    newMessage.classList.add('message');
    if (isUser) {
        newMessage.classList.add('user-message');
    } else {
        newMessage.classList.add('other-message');
    }
    newMessage.textContent = message;

    // Добавляем временную метку
    const timestamp = document.createElement('span');
    timestamp.classList.add('timestamp');
    timestamp.textContent = new Date().toLocaleTimeString();
    newMessage.appendChild(timestamp);

    chatMain.appendChild(newMessage);

    try{
        if(isUser !== false)
            sendMessage(message,"chat");
    }catch (e) {

    }


}

// Пример добавления сообщения от вас
addMessage("ку", false);

// Обработчик события для кнопки отправки
sendButton.addEventListener('click', () => {
    const message = messageInput.value;
    if (message.trim() !== '') {
        addMessage(message, true);
        messageInput.value
            = '';
    }
});

messageInput.addEventListener('keydown', (event) => {
    if (event.key === 'Enter' && document.activeElement === messageInput) {
        const message = messageInput.value;
        if (message.trim() !== '') {
            addMessage(message, true);
            messageInput.value = '';
        }
    }
});