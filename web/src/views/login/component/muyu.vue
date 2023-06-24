<template>
  <div>
    <div class="w-f-container">
      <div :class="state.muyuTextClass">叼毛 -1</div>
      <img @click="clickMuyu" :class="state.muyuImgClass" :src="muyuImg"/>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {reactive} from "vue";
//mp3
import muyuMp3 from "/@/assets/music/muyu.mp3"
//muyuImg
import muyuImg from "/@/assets/icon/wooden_fish.png"

const state = reactive({
  muyuImgClass: 'w-f-c-img',
  muyuTextClass: 'w-f-c-text'
})

//点击播放
function clickMuyu() {
  let audio = new Audio()
  audio.src = muyuMp3
  audio.play();
  state.muyuImgClass = 'w-f-c-img w-f-c-i-size'
  state.muyuTextClass = 'w-f-c-text w-f-c-t-size'
  setTimeout(() => {
    state.muyuImgClass = 'w-f-c-img'
    state.muyuTextClass = 'w-f-c-text'
  }, 300)
}

</script>

<style scoped>

.w-f-container {
  width: 80px;
  position: absolute;
  left: 50%;
  top: 88%;
  transform: translate(-50%, -50%); /* 50%为自身尺寸的一半 */
}

.w-f-c-img {
  width: 100%;
  height: 100%;
}

.w-f-c-text {
  width: 100%;
  font-size: 18px;
  text-align: center;
  margin-bottom: 30px;
  /*position: absolute;*/
  animation: wooden-fish-text-move 1s forwards;
}

@keyframes wooden-fish-text-move {
  0% {
    top: -50px;
    opacity: 1;
  }
  100% {
    top: -110px;
    opacity: 0;
  }
}

.w-f-c-i-size {
  animation: wooden-fish-size 0.3s;
}

.w-f-c-t-size {
  animation: wooden-fish-text-move 0.3s;
}

@keyframes wooden-fish-size {
  0% {
    transform: scale(1);
  }
  25% {
    transform: scale(1.5);
  }
  50% {
    transform: scale(2);
  }
  75% {
    transform: scale(1.5);
  }
  100% {
    transform: scale(1);
  }
}

</style>