<template>
  <div id="app" class="container">
    <div class="row">
      <a href="javascript:;" @click="isEditList = !isEditList">Chinh sua danh sach</a>
    </div>
    <div v-show="isEditList">
      <div class="row">
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
      <h4>Tổng số người tham gia: {{mems.length}}</h4>
      <div style="margin-right: auto"></div>
      <a href="javascript:;" @click="onCreateTeamRandom" class="btn btn-secondary">Tạo ngẫu nhiên</a>
    </div>
    <div class="row" style="margin-top: 20px">
      <a href="javascript:;" @click="back" class="btn btn-secondary">Back</a>
      <a href="javascript:;" style="margin-left: 10px" @click="next" class="btn btn-secondary">Next</a>
      <div style="margin-right: auto"></div>
      <div class="form-check btn">
        <input type="checkbox" class="form-check-input" id="exampleCheck1" v-model="dispPresenter">
        <label class="form-check-label" for="exampleCheck1">Ẩn danh sách nhận quà</label>
      </div>
    </div>
    <div class="row team_area">
      <div class="col-5">
        <h4>Bên tặng quà</h4>
        <ul class="list-group">
          <li class="list-group-item"
              :class="{'active': idx === active, 'hide':  (idx < active || idx > active)&& dispPresenter}"
              v-for="(i,idx) in presenters" :key="idx">
            {{i.name}}
          </li>
        </ul>
      </div>
      <div style="margin-right: auto"></div>
      <div class="col-5">
        <h4>Bên nhận quà</h4>
        <ul class="list-group">
          <li class="list-group-item"
              :class="{'active': idx === active, 'hide':   (idx < active || idx > active) && dispPresenter }"
              v-for="(i,idx) in receivers" :key="idx">
            {{i.name}}
          </li>
        </ul>
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
      };
    },
    computed: {},
    watch: {
      mems() {
        window.localStorage.setItem('users', JSON.stringify(this.mems));
        this.randomTeam()
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
      speak() {
        let send = this.presenters[this.active];
        let rect = this.receivers[this.active];
        if (!send || !rect) {
          return;
        }
        this.queue_sounds([
          new Audio(require('../voice/' + send.voice_send + '.mp3')),
          new Audio(require('../voice/' + rect.voice_rec + '.mp3'))],
        );
      },
      randomTeam() {
        this.presenters = this.mems.concat().sort(function() {
          return 0.5 - Math.random();
        });
        let temp = this.mems.slice();
        let rs = [];
        while (temp.length > 0) {
          rs.push(temp.splice(this.getRandUniqueRec(rs.length, temp), 1)[0]);
        }
        this.receivers = rs;
      },
      onCreateTeamRandom() {
        if (confirm('Bạn có chắc muốn random lại?')) {
          this.randomTeam();
        }
      },
      back() {
        if (this.active > -1) {
          this.active--;
        }
        this.speak();
      },
      next() {
        if (this.active < this.mems.length) {
          this.active++;
        }
        this.speak();
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
      this.randomTeam();
      // window.addEventListener('scroll', this.handleDebouncedScroll);
    },
    beforeDestroy() {
      // I switched the example from `destroyed` to `beforeDestroy`
      // to exercise your mind a bit. This lifecycle method works too.
      // window.removeEventListener('scroll', this.handleDebouncedScroll);
    },
  };
</script>

<style scoped>
  .hide {
    display: none !important;
  }

  #app {
    margin-top: 20px;
    padding: 20px;
  }

  .team_area {
    margin-top: 20px;
  }

</style>
