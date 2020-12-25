<template>
  <div class="about">
    <button class="btn btn-success" style="margin-top: 20px; width: 80%;" @click="speak('1')">
      Tiep tuc
    </button>
    <ul class="nav events">
      <li v-for="(i, idx) in sounds" :key="idx" class="nav-item">
        <a href="javascript:;" class="btn " @click="speak(i.voice)">{{ i.name }}</a>
      </li>
    </ul>
  </div>
</template>

<script>
/* eslint-disable */
import {db} from '../fb'
export default {
  data(){
    return {
      doc: {},
      btns: [
        {}
      ],
      sounds: [
        {'name': '👏', voice: '2'},
        {'name': '🕐', voice: '3'},
        {'name': '👌', voice: '4'},
        {'name': '👎', voice: '5'},
        {'name': '👍', voice: '6'},
      ],
    }
  },
  methods: {
     async speak(val){
      await db.collection('command').doc('noel').update({cmd: val, time: new Date().getTime()})
    }
  }
}
</script>


<style scoped>
.li{
  width: 100%;
  margin: 50px;
}
.li .btn{
  width: 100%;
}

.events {
  position: fixed;
  bottom: 20px;
}

.events li {
  margin: 0 10px;
  background: white;
  border-radius: 100%;
}

.events li a {
  font-size: 30px !important;
}
</style>