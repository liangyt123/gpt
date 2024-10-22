const API_BASE_URL = "http://192.168.116.35:80"; // 根据你的后端地址配置
let token = "";
let click_lock = false;
let playerInfo = {};
let current_step = 0;

// 获取玩家信息
function getPlayerInfo() {
  fetch(`${API_BASE_URL}/api/player`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ token }),
  })
    .then((response) => response.json())
    .then((data) => {
      current_step = data.current_step;
      // 更新玩家信息和故事背景
      playerInfo = Object.assign(playerInfo, data);
      document.getElementById("playerInfo").innerHTML = `
                <p class="playerInfo_name">名称: <span class="playerInfo_name_span">${playerInfo.name}</span></p>
                <div class="playerInfo_value">${data.territory}<div class="playerInfo_value_text">爱戴值</div></div>
                <p class="playerInfo_time">当前回合: <span class="playerInfo_time_span">${data.current_step}</span></p>
            `;
      document.getElementById("storyBackground").innerHTML = `
                <p><span class="storyBackground_firstText">背景:</span><span id="text-to-display2">  </span></p>
            `;
      document.getElementById("storyPlot").innerHTML = `
                <p id="storyPlot_P"><span  class="storyBackground_firstText">剧情:</span> <span id="text-to-display"></span></p>
            `;

      token = data.token; // 保存 token
      // 更新按钮文本

      console.log("aaa:", data);
      document.getElementById("choiceA").hidden = false;
      document.getElementById("choiceA").innerText = data.choice_a;
      document.getElementById("choiceB").innerText = data.choice_b;

      document.getElementById("nameDialog").hidden = true;

      const displayElement = document.getElementById("text-to-display");
      const displayElement2 = document.getElementById("text-to-display2");
      const typingContainer = document.getElementById("storyPlot_P");
      const typingContainer2 = document.getElementById("storyBackground");
      let index = 0;
      let index2 = 0;
      let time = null;
      let time2 = null;
      displayElement.textContent = "";
      displayElement2.textContent = "";
      function typeText() {
        if (index < data.background.length) {
          displayElement2.textContent += data.background[index];
          index++;
          time2 = setTimeout(typeText, 70); // 调整时间间隔以控制打字速度
          if (typingContainer2.scrollHeight > typingContainer2.clientHeight) {
            typingContainer2.scrollTop =
              typingContainer2.scrollHeight - typingContainer2.clientHeight;
          }
        } else if (index2 < data.story.length) {
          displayElement.textContent += data.story[index2];
          index2++;
          time = setTimeout(typeText, 70); // 调整时间间隔以控制打字速度
          if (typingContainer.scrollHeight > typingContainer.clientHeight) {
            typingContainer.scrollTop =
              typingContainer.scrollHeight - typingContainer.clientHeight;
          }
        }
      }
      if (time) clearTimeout(time);
      typeText();
    });
}

// 提交玩家选择
function makeChoice(choice) {
  if (click_lock) return;
  click_lock = true;
  setButtonStatus(false);
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
        logContent.innerHTML += `<p id="${"log_" + current_step}" >${
          data.result
        }</p>`;
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
          let storyPlot = document.getElementById("storyPlot");
          storyPlot.style.background = `url(${imgUrl}) no-repeat center center`;
          storyPlot.style.backgroundSize = "100% 100%";
          let log = document.getElementById("log_" + (current_step - 1));
          const img = document.createElement("img");
          img.src = imgUrl;
          img.style.width = "200px";
          img.style.height = "200px";
          img.style.margin = "0 auto";
          img.style.display = "block";
          log.parentNode.insertBefore(img, log.nextSibling);
          //   let content2 =
          //     `
          //             <p>达芬奇用画笔记录了这一刻：</p>
          //             <img class="storyPlot_img" src="` +
          //     imgUrl +
          //     `></iframe>
          //         `;
          //   console.log("content:", content2);
          //   document.getElementById("storyPlot").appendChild(content2);
        }
        click_lock = false;
        setButtonStatus(true);
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
    click_lock = false;
    setButtonStatus(true);
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
    let storyPlot = document.getElementById("storyPlot");
    storyPlot.style.background = `url(${"/static/img/dialog_back.jpeg"}) no-repeat center center`;
    storyPlot.style.backgroundSize = "100% 100%";
    setButtonStatus(true);
  } else {
    alert("请输入玩家名称");
    return;
  }
}

function setButtonStatus(show) {
  document.getElementById("showButton").style.display = show ? "block" : "none";
  document.getElementById("showLoading").style.display = show
    ? "none"
    : "block";
}
