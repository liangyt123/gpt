const API_BASE_URL = "http://192.168.116.35:80"; // 根据你的后端地址配置
let token = "";
let playerInfo = {};

// 获取玩家信息
function getPlayerInfo() {
  fetch(`${API_BASE_URL}/api/player`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ token }),
  })
    .then((response) => response.json())
    .then((data) => {
      // 更新玩家信息和故事背景
      playerInfo = Object.assign(playerInfo, data);
      document.getElementById("playerInfo").innerHTML = `
                <p class="playerInfo_name">名称: <span class="playerInfo_name_span">${playerInfo.name}</span></p>
                <div class="playerInfo_value">${data.territory}<div class="playerInfo_value_text">爱戴值</div></div>
                <p class="playerInfo_time">当前回合: <span class="playerInfo_time_span">${data.current_step}</span></p>
            `;
      document.getElementById("storyBackground").innerHTML = `
                <p><span class="storyBackground_firstText">背景:</span> ${data.background}</p>
            `;
      document.getElementById("storyPlot").innerHTML = `
                <p><span class="storyBackground_firstText">剧情:</span> ${data.story}</p>
            `;

      token = data.token; // 保存 token
      // 更新按钮文本

      console.log("aaa:", data);
      document.getElementById("choiceA").hidden = false;
      document.getElementById("choiceA").innerText = data.choice_a;
      document.getElementById("choiceB").innerText = data.choice_b;

      document.getElementById("nameDialog").hidden = true;
    });
}

// 提交玩家选择
function makeChoice(choice) {
  fetch(`${API_BASE_URL}/api/choose`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ choice, token }),
  })
    .then((response) => response.json())
    .then((data) => {
      // 更新日志框内容
      let logContent = document.getElementById("logContent");
      console.log(data);
      // 如果有结果，添加到日志框 且不是 undefined
      if (data.result) {
        logContent.innerHTML += `<p>${data.result}</p>`;
      }
      document.getElementById("miniGame").innerHTML = ``;
      // 滚动到日志底部
      logContent.scrollTop = logContent.scrollHeight;

      getPlayerInfo(); // 重新加载玩家信息
    });
  // 增加随机数 随机展示小游戏
  let random = Math.floor(Math.random() * 10);

  let story = playerInfo.story;
  console.log("story:", story, "random:", random);
  if (random % 10 >= 0) {
    fetch(`${API_BASE_URL}/api/generate`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ story, token }),
    })
      .then((response) => response.json())
      .then((data) => {
        console.log("data1:", data);
        let imgUrl = data.current_choice.img_url;
        // 如果 img_url 含有链接则展示图片
        if (imgUrl == undefined) {
          return;
        }
        if (imgUrl.includes("http")) {
          let content2 =
            `
                    <p>达芬奇用画笔记录了这一刻：</p>
                    <img src="` +
            imgUrl +
            `" width="375px" height="375px"></iframe>
                `;
          console.log("content:", content2);
          document.getElementById("miniGame").innerHTML = content2;
        }
      });
    // 如果是 %10 == 1 强制开始游戏
  } else if (true || random % 10 == 1) {
    document.getElementById("miniGame").innerHTML = `
            <p>小游戏: 连连看</p>
            <iframe src="static/minigame/link-game/index.html" width="100%" height="375px"></iframe>
        `;
    // 隐藏按钮
    // document.getElementById("choiceA").innerText = data.choice_a;
    document.getElementById("choiceA").hidden = true;
    document.getElementById("choiceB").hidden = true;
    // document.getElementById("choiceB").innerText = data.choice_b;
  }
}

// 页面加载时获取玩家信息
// document.addEventListener("DOMContentLoaded", getPlayerInfo);

function startGame() {
  let name = document.getElementById("playerName").value;
  console.log(name);
  if (name) {
    playerInfo.name = name;
    document.getElementById("startGameButton").style.display = "none";
    document.getElementById("loading").style.display = "block";
    document.getElementById("loading_text").style.display = "block";
    getPlayerInfo();
  } else {
    alert("请输入玩家名称");
    return;
  }
}
