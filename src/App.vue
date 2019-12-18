<template>
  <div id="app" class="container">
    <div class="row" style="margin-bottom: 20px">
      <a href="javascript:;" @click="isEditList = !isEditList">{{menuTitle}}

      </a>

    </div>
    <div v-show="isEditList">
      <div class="row">
        <p>Mỗi lần chỉnh sửa danh sách tham dự là danh sách người tặng quà và nhận quà sẽ tự động sắp xếp ngẫu nhiên</p>
        <div style="margin-right: auto"></div>
        <a href="javascript:;" class="btn btn-danger" @click="reset">Reset</a>
      </div>
      <div class="row">
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
    <div class="row">
      <h4>Tổng số người tham gia: {{mems.length}}/{{users.length}}</h4>
      <div style="margin-right: auto"></div>
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
        <div class="btn_c" @click="onGetSender" :class="{'disabled': flg !== false}">
          <h4>Tặng quà</h4>
          <span v-if="flg === false">👆🏻</span>
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
        <div class="btn_c" @click="onGetReceiver" :class="{'disabled': flg !== true}">
          <h4>Nhận quà</h4>
          <span v-if="flg === true">👆🏻</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import List from './list.json';

  export default {
    name: 'app',
    data() {
      return {
        users: [],
        isEditList: false,
        presenters: [],
        receivers: [],
        active: -1,
        dispPresenter: true,
        mems: [],
        sender: {},
        receiver: {},
        flg: false,
      };
    },
    computed: {
      menuTitle() {
        return this.isEditList ? 'Đóng chỉnh sửa danh sách tham dự' : 'Mở chỉnh sửa danh sách tham dự';
      },
    },
    watch: {
      mems() {
        window.localStorage.setItem('users', JSON.stringify(this.mems));
        this.presenters = this.mems.slice();
        this.receivers = this.mems.slice();
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
        if (confirm('Chắc chắn chưa?')) {
          this.sender = this.presenters.splice(Math.floor(Math.random() * this.presenters.length), 1)[0];
          this.flg = !this.flg;
          this.receiver = {};
          const t = this;
          setTimeout(function() {
            t.speak(t.sender.voice_send);
          }, 2000);
        }

      },
      onGetReceiver() {
        if (confirm('Chắc chắn chưa?')) {
          this.receiver = this.receivers.splice(Math.floor(Math.random() * this.presenters.length), 1)[0];
          this.flg = !this.flg;
          const t = this;
          setTimeout(function() {
            t.speak(t.receiver.voice_rec);
          }, 2000);

        }
      },
      delay(time) {
        setTimeout(function() {
          //your code to be executed after 1 second
        }, time);
      },
      getRandUniqueRec(i, arr) {
        let idx = -1;
        if (arr.length == 1) {
          return 0;
        }
        do {
          idx = Math.floor(Math.random() * arr.length);
        } while (arr[idx] == this.presenters[i]);
        return idx;
      },
      reset() {
        if (confirm('Bạn có chắc chắn muốn danh sách member về lại như ban đầu không?')) {
          this.mems = this.users.slice();
        }
      },
    },
    created() {
      this.queue_sounds([new Audio(require('../voice/wellcome.mp3'))]);
      this.users = List;
      if (window.localStorage.getItem('users')) {
        this.mems = JSON.parse(window.localStorage.getItem('users'));
      } else {
        this.mems = this.users.slice();
      }
      this.presenters = this.mems.slice();
      this.receivers = this.mems.slice();
    },
    beforeDestroy() {
    },
  };
</script>

<style scoped>
  .disabled {
    pointer-events: none;
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
    margin-top: 40px;
    user-select: none;
    -webkit-user-select: none;
  }

  .btn_c {
    margin-top: 100px;
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
