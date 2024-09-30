const API_BASE_URL = 'http://localhost:8080'; // 后端 API 地址

let currentTurn = 1; // 当前回合

// script.js
window.onload = function () {
    fetchPlayerInfo(); // 页面加载时获取玩家信息
    // 获取按钮元素
    const drawCardBtn = document.getElementById('draw-card-btn');
    const triggerEventBtn = document.getElementById('trigger-event-btn');
    const useCardBtn = document.getElementById('use-card-btn');
    const nextTurnBtn = document.getElementById('next-turn-btn');

    // 为按钮绑定点击事件
    drawCardBtn.addEventListener('click', function() {
        drawCard();
        // 在这里执行抽取卡牌的逻辑
    });

    triggerEventBtn.addEventListener('click', function() {
        triggerEvent();
        // 在这里执行触发事件的逻辑
    });

    useCardBtn.addEventListener('click', function() {
        useCard();
        // 在这里执行使用卡牌的逻辑
    });

    nextTurnBtn.addEventListener('click', function() {
        nextTurn();
        // 在这里执行跳过回合的逻辑
    });
};

// 初始化玩家信息
function fetchPlayerInfo() {
    fetch(`${API_BASE_URL}/player`)
        .then(response => response.json())
        .then(data => {
            console.log('player info:', data);
            document.getElementById('player-name').innerText = `玩家名称: ${data.Name}`;
            document.getElementById('player-resources').innerText = `资源: ${data.Resources}`;
            document.getElementById('player-army').innerText = `军队: ${data.Army}`;
            document.getElementById('player-turn').innerText = `当前回合: ${currentTurn}`;
        })
        .catch(error => console.error('Error fetching player info:', error));
}

// 抽取卡牌
function drawCard() {
    fetch(`${API_BASE_URL}/draw-card`)
        .then(response => response.json())
        .then(card => {
            console.log('card:', card);
            const cardContainer = document.getElementById('card-container');
            const cardElement = document.createElement('div');
            cardElement.classList.add('card');
            cardElement.innerHTML = `<h3>${card.Name}</h3><p>力量: ${card.Power}</p><p>技能: ${card.Skill}</p>`;
            cardContainer.appendChild(cardElement);
            logAction(`抽取了一张卡牌: ${card.Name}`);
        })
        .catch(error => console.error('Error drawing card:', error));
}

// 使用卡牌
function useCard() {
    // 假设每次使用的都是第一张卡牌
    const cardContainer = document.getElementById('card-container');
    if (cardContainer.children.length > 0) {
        const cardElement = cardContainer.children[0];
        const cardName = cardElement.querySelector('h3').innerText;
        cardContainer.removeChild(cardElement); // 移除已使用的卡牌
        logAction(`使用了卡牌: ${cardName}`);
        // 更新玩家信息，假设卡牌会影响资源
        updateResourcesAfterCardUse();
    } else {
        logAction('没有可用的卡牌');
    }
}

// 更新资源
function updateResourcesAfterCardUse() {
    // 假设每次使用卡牌减少 10 资源
    fetch(`${API_BASE_URL}/use-card`, { method: 'POST' })
        .then(() => {
            fetchPlayerInfo(); // 更新玩家信息
        })
        .catch(error => console.error('Error using card:', error));
}

// 触发事件
function triggerEvent() {
    fetch(`${API_BASE_URL}/trigger-event`)
        .then(response => response.json())
        .then(event => {
            console.log('event:', event);
            document.getElementById('event-description').innerText = `当前事件: ${event.Description}`;
            fetchPlayerInfo(); // 更新玩家信息
            logAction(`触发了事件: ${event.Description}`);
        })
        .catch(error => console.error('Error triggering event:', error));
}

// 跳过回合
function nextTurn() {
    currentTurn++;
    document.getElementById('player-turn').innerText = `当前回合: ${currentTurn}`;
    logAction(`跳过了回合 ${currentTurn}`);
}

// 记录操作日志
function logAction(action) {
    const logContainer = document.getElementById('log-container');
    const logElement = document.createElement('p');
    logElement.innerText = action;
    logContainer.appendChild(logElement);
}

