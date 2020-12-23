<template>
  <div id="app" class="container">
    <div class="row" style="margin-bottom: 20px">
      <a href="javascript:;" @click="isEditList = !isEditList">{{menuTitle}}
      </a>
    </div>
    <div v-show="isEditList">
      <div class="" style="display: flex">
        <p>Chỉnh sửa danh sách tham dự sẽ reset danh sách người tặng quà và nhận quà lại ban đầu</p>
        <div style="margin-right: auto"></div>
        <div><a href="javascript:;" class="btn btn-danger" @click="reset">Reset</a></div>
      </div>
      <div class="row">
        <div style="overflow-y: auto; max-height: 400px; border: 1px solid #ccc; padding: 20px 0">
          <ul class="list-group">
            <li class="list-group-item" v-for="(i, idx) in users" :key="idx">
              <div class="form-check btn">
                <input class="form-check-input" type="checkbox" :id="idx + '-checkuser'" v-model="mems" :value="i">
                <label class="form-check-label" :for="idx + '-checkuser'">{{i.name}}</label>
              </div>
            </li>
          </ul>
        </div>
      </div>
    </div>
    <div style="margin-top: 20px" class="row">
      <h4>Tổng số người tham gia: {{mems.length}}/{{users.length}}</h4>
    </div>
    <div class="row">
      <p>Còn lại: {{presenters.length}} người tặng vs {{receivers.length}} người nhận</p>
    </div>
    <div class="row team_area">
      <div class="col-6">
        <ul class="list-group">
          <transition name="slide-fade" mode="out-in">
            <template v-if="sender.name">
              <li class="list-group-item typewriter active" :key="sender.name">
                {{sender.name}}
              </li>
            </template>
            <template v-else>
              <li class="list-group-item finder">
                🧐
              </li>
            </template>
          </transition>
        </ul>
        <div class="btn_c" @click="onGetSender" :class="{'disabled': flg !== false ||  presenters.length===0}">
          <h4>Tặng quà</h4>
          <span v-if="flg === false && presenters.length">👆🏻</span>
        </div>
      </div>
      <div style="margin-right: auto"></div>
      <div class="col-6">
        <ul class="list-group">
          <transition name="slide-fade" mode="out-in">
            <template v-if="receiver.name">
              <li class="list-group-item typewriter active" :key="receiver.name">
                {{receiver.name}}
              </li>
            </template>
            <template v-else>
              <li class="list-group-item finder">
                🧐
              </li>
            </template>
          </transition>
        </ul>
        <div class="btn_c" @click="onGetReceiver" :class="{'disabled': flg !== true || receivers.length === 0}">
          <h4>Nhận quà</h4>
          <span v-if="flg === true && receivers.length">👆🏻</span>
        </div>
      </div>
    </div>
    <div class="row" v-show="!isEditList">
      <ul class="nav events">
        <li v-for="(i, idx) in sounds" :key="idx" class="nav-item">
          <a href="javascript:;" class="btn " @click="speak(i.voice)">{{i.name}}</a>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
  import List from './list.json';

  export default {
    name: 'app',
    data() {
      return {
        sounds: [
          {'name': '👏', voice: 'clap'},
          {'name': '🕐', voice: 'timeout'},
          {'name': '👌', voice: 'correct'},
          {'name': '👎', voice: 'wrong'},
          {'name': '👍', voice: 'win'},
        ],
        users: [],
        isEditList: false,
        presenters: [],
        receivers: [],
        active: -1,
        mems: [],
        sender: {},
        receiver: {},
        flg: false,
        errors: {},
        isLoaded: false,
      };
    },
    computed: {
      menuTitle() {
        return this.isEditList ? 'Đóng chỉnh sửa danh sách tham dự' : 'Mở chỉnh sửa danh sách tham dự';
      },
    },
    watch: {
      mems() {
        if (this.isLoaded) {
          window.localStorage.setItem('users', JSON.stringify(this.mems));
          this.presenters = this.mems.slice();
          this.receivers = this.mems.slice();
          this.sender = {};
          this.receiver = {};
          this.flg = false;
        } else {
          this.isLoaded = true;
        }
      },
      sender(val) {
        window.localStorage.setItem('sender', JSON.stringify(val));
      },
      receiver(val) {
        window.localStorage.setItem('receiver', JSON.stringify(val));
      },
      presenters(val) {
        window.localStorage.setItem('presenters', JSON.stringify(val));
        window.localStorage.setItem('sender', JSON.stringify(this.sender));
      },
      receivers(val) {
        window.localStorage.setItem('receivers', JSON.stringify(val));
        window.localStorage.setItem('receiver', JSON.stringify(this.receiver));
      },
      flg(val) {
        window.localStorage.setItem('flg', JSON.stringify(val));
      },
    },
    methods: {
      play(audio, callback) {
        audio.play();
        if (callback) {
          audio.onended = callback;
        }
      },

      queue_sounds(sounds) {
        var index = 0;
        const t = this;

        function recursive_play() {
          if (index + 1 === sounds.length) {
            t.play(sounds[index], null);
          } else {
            t.play(sounds[index], function() {
              index++;
              recursive_play();
            });
          }
        }

        recursive_play();
      },
      speak(file) {
        if (!file) {
          return;
        }
        this.queue_sounds([new Audio(require('../voice/' + file + '.mp3'))]);
      },
      onGetSender() {
          this.sender = this.presenters.splice(Math.floor(Math.random() * this.presenters.length), 1)[0];
          this.flg = !this.flg;
          this.receiver = {};
          const t = this;
          setTimeout(function() {
            t.speak(t.sender.voice_send);
          }, 2000);

      },
      onGetReceiver() {
          let idx = this.getRandUniqueRec(this.receivers);
          this.receiver = this.receivers.splice(idx, 1)[0];
          this.flg = !this.flg;
          const t = this;
          setTimeout(function() {
            t.speak(t.receiver.voice_rec);
          }, 2000);
      },
      delay(time) {
        setTimeout(function() {
          //your code to be executed after 1 second
        }, time);
      },
      getRandUniqueRec(arr) {
        let idx = -1;
        let len = arr.length;
        if (len === 1) {
          return 0;
        }
        if (len === 2) {
          return arr[0].name.localeCompare(this.sender.name) !== 0
          && arr[1].name.localeCompare(this.presenters[0].name) === 0 ? 1 : 0;
        }
        do {
          idx = Math.floor(Math.random() * len);
        } while (arr[idx].name.localeCompare(this.sender.name) === 0);

        return idx;
      },
      reset() {
        if (confirm('Bạn có chắc chắn muốn danh sách member về lại như ban đầu không?')) {
          this.mems = this.users.slice();
        }
      },
    },
    mounted() {
    },
    created() {
      window.onload = function() {
        var audio = document.createElement('AUDIO');
        document.body.appendChild(audio);
        audio.play();
      };
      this.users = List;
      if (window.localStorage.getItem('users')) {
        this.mems = JSON.parse(window.localStorage.getItem('users'));
      } else {
        this.mems = this.users.slice();
      }
      if (window.localStorage.getItem('presenters')) {
        this.presenters = JSON.parse(window.localStorage.getItem('presenters'));
      }
      if (window.localStorage.getItem('receivers')) {
        this.receivers = JSON.parse(window.localStorage.getItem('receivers'));
      }
      if (window.localStorage.getItem('sender')) {
        this.sender = JSON.parse(window.localStorage.getItem('sender'));
      }
      if (window.localStorage.getItem('receiver')) {
        this.receiver = JSON.parse(window.localStorage.getItem('receiver'));
      }
      if (window.localStorage.getItem('flg')) {
        this.flg = JSON.parse(window.localStorage.getItem('flg'));
      }
    },
    beforeDestroy() {
    },
  };
</script>

<style scoped>
  .disabled {
    pointer-events: none;
    border: none !important;
  }

  .finder {
    font-size: 32px;
    text-align: center;
    padding: 0;
  }

  .slide-fade-enter-active {
    transition: all 1s ease;
  }

  .slide-fade-leave-active {
    transition: all 2s cubic-bezier(1.0, 0.5, 0.8, 1.0);
  }

  .slide-fade-enter, .slide-fade-leave-to
    /* .slide-fade-leave-active for <2.1.8 */
  {
    transform: translateX(10px);
    opacity: 0;
  }

  .hide {
    display: none !important;
  }

  .list-group {
    user-select: none;
    -webkit-user-select: none;
  }

  .events {
    position: fixed;
    bottom: 20px;
  }

  .events li {
    border: 1px solid #ccc;
  }

  .events li a {
    font-size: 30px !important;
  }

  .btn_c {
    border: 1px solid #ccc;
    margin-top: 100px;
    padding-top: 4px;
    cursor: pointer;
    display: flex;
    position: relative;
    flex-direction: column;
    align-items: center;
    user-select: none;
  }

  .btn_c span {
    font-size: 30px;
    position: absolute;
  }

  .btn_c span {
    -webkit-animation-name: moveit; /* Safari 4.0 - 8.0 */
    -webkit-animation-duration: 1.5s; /* Safari 4.0 - 8.0 */
    animation-name: moveit;
    animation-duration: 1.5s;
    animation-iteration-count: infinite;
  }

  /* Safari 4.0 - 8.0 */
  @-webkit-keyframes moveit {
    from {
      top: 15px
    }
    to {
      top: 30px
    }
  }

  /* Standard syntax */
  @keyframes moveit {
    from {
      top: 15px
    }
    to {
      top: 30px
    }
  }

  #app {
    margin-top: 20px;
    padding: 20px;
  }

  .team_area {
    margin-top: 20px;
  }

  /*.typewriter {*/
  /*  overflow: hidden; !* Ensures the content is not revealed until the animation *!*/
  /*  border-right: .15em solid orange; !* The typwriter cursor *!*/
  /*  white-space: nowrap; !* Keeps the content on a single line *!*/
  /*  margin: 0 auto; !* Gives that scrolling effect as the typing happens *!*/
  /*  letter-spacing: .15em; !* Adjust as needed *!*/
  /*  animation: typing 10s steps(40, end),*/
  /*  blink-caret .75s step-end infinite;*/
  /*}*/

  /*!* The typing effect *!*/
  /*@keyframes typing {*/
  /*  from {*/
  /*    width: 0*/
  /*  }*/
  /*  to {*/
  /*    width: 100%*/
  /*  }*/
  /*}*/

  /* The typewriter cursor effect */
  @keyframes blink-caret {
    from, to {
      border-color: transparent
    }
    50% {
      border-color: orange;
    }
  }
</style>
