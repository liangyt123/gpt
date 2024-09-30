const API_BASE_URL = 'http://localhost:80'; // 根据你的后端地址配置
let token = "";

// 获取玩家信息
function getPlayerInfo() {
    fetch(`${API_BASE_URL}/api/player`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ token })
    })
        .then(response => response.json())
        .then(data => {
            // 更新玩家信息和故事背景
            document.getElementById('playerInfo').innerHTML = `
                <p>领土: ${data.territory}</p>
                <p>当前步骤: ${data.current_step}</p>
            `;
            document.getElementById('storyBackground').innerHTML = `
                <p>故事背景: ${data.story}</p>
            `;
            console.log('player info:', data);
            token = data.token;  // 保存 token
            // 更新按钮文本
            document.getElementById('choiceA').innerText = data.choice_a;
            document.getElementById('choiceB').innerText = data.choice_b;
        });
}

// 提交玩家选择
function makeChoice(choice) {
    fetch(`${API_BASE_URL}/api/choose`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ choice,token })
    })
    .then(response => response.json())
    .then((data) => {
        // 更新日志框内容
        let logContent = document.getElementById('logContent');
        console.log(data);
        // 如果有结果，添加到日志框 且不是 undefined
        if (data.result) {
            logContent.innerHTML += `<p>${data.result}</p>`;
        }
        
        // 滚动到日志底部
        logContent.scrollTop = logContent.scrollHeight;

        getPlayerInfo();  // 重新加载玩家信息
    });
}

// 页面加载时获取玩家信息
document.addEventListener('DOMContentLoaded', getPlayerInfo);
