const API_BASE_URL = 'http://127.0.0.1:80'; // 根据你的后端地址配置
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
                <p>爱戴值: ${data.territory}</p>
                <p>当前回合: ${data.current_step}</p>
                <p>名称: ${data.token}</p>
            `;
            document.getElementById('storyBackground').innerHTML = `
                <p>背景: ${data.background}</p>
            `;
            document.getElementById('storyPlot').innerHTML = `
                <p>剧情: ${data.story}</p>
            `;
          
            token = data.token;  // 保存 token
            // 更新按钮文本

            console.log('aaa:', data);
            if (data.mini_game != "") {
                  // 如果是小游戏 强制开始游戏
                if (data.mini_game == "连连看") {
                    document.getElementById('miniGame').innerHTML = `
                        <p>小游戏: ${data.mini_game}</p>
                        <iframe src="static/minigame/link-game/index.html" width="100%" height="660px"></iframe>
                    `;
                }
                if (data.mini_game == "打砖块") {
                    document.getElementById('miniGame').innerHTML = `
                        <p>小游戏: ${data.mini_game}</p>
                        <iframe src="static/minigame/h5-game-blockBreaker/index.html" width="100%" height="660px"></iframe>
                    `;
                }

                if (data.mini_game == "合成大西瓜") {
                    document.getElementById('miniGame').innerHTML = `
                        <p>小游戏: ${data.mini_game}</p>
                        <iframe src="static/minigame/bigwatermelon/index.html"  height="660px"></iframe>
                    `;
                }
                document.getElementById('choiceA').innerText = data.choice_a;
                document.getElementById('choiceA').hidden = true;    
                document.getElementById('choiceB').innerText = data.choice_b;
            }else {
                document.getElementById('choiceA').hidden = false;   
                document.getElementById('choiceA').innerText = data.choice_a;            
                document.getElementById('choiceB').innerText = data.choice_b;
            }
    
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
        document.getElementById('miniGame').innerHTML = ``;
        // 滚动到日志底部
        logContent.scrollTop = logContent.scrollHeight;

        getPlayerInfo();  // 重新加载玩家信息
    });
}

// 页面加载时获取玩家信息
document.addEventListener('DOMContentLoaded', getPlayerInfo);
