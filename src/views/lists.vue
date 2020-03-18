<template>
  <div id="listContent">
    <div id="list">
      <div class="header">
        <h1>
          {{id}}
        </h1>
        <div class="edit evocationModal">
          <svg viewBox="0 0 24 24">
            <path d="M13.5,17c0,0.8-0.7,1.5-1.5,1.5s-1.5-0.7-1.5-1.5s0.7-1.5,1.5-1.5S13.5,16.2,13.5,17z M13.5,12c0,0.8-0.7,1.5-1.5,1.5 s-1.5-0.7-1.5-1.5s0.7-1.5,1.5-1.5S13.5,11.2,13.5,12z M13.5,7c0,0.8-0.7,1.5-1.5,1.5S10.5,7.8,10.5,7s0.7-1.5,1.5-1.5 S13.5,6.2,13.5,7z"></path>
          </svg>
        </div>
      </div>
      <div class="filter">
        <div class="main">
          <div class="selector-number">
            <p>5 篇笔记</p>
          </div>
          <div class="sort">
            <svg viewBox="0 0 1024 1024">
              <path d="M611.983,162.025H61.999C34.384,162.025,12,184.409,12,212.023c0,27.614,22.384,49.999,49.999,49.999 h549.985c27.614,0,49.999-22.384,49.999-49.999C661.982,184.409,639.598,162.025,611.983,162.025z M611.983,462.016h-99.997 c-27.614,0-49.999,22.384-49.999,49.999c0,27.614,22.384,49.999,49.999,49.999h99.997c27.614,0,49.999-22.384,49.999-49.999 C661.982,484.401,639.598,462.016,611.983,462.016z M61.999,562.014h249.993c27.614,0,49.999-22.384,49.999-49.999 c0-27.614-22.384-49.999-49.999-49.999H61.999C34.384,462.016,12,484.401,12,512.015C12,539.629,34.384,562.014,61.999,562.014z M611.983,762.008H61.999C34.384,762.008,12,784.392,12,812.007s22.384,49.999,49.999,49.999h549.985 c27.614,0,49.999-22.384,49.999-49.999S639.598,762.008,611.983,762.008z M997.533,637.187c-19.734-19.799-51.729-19.264-70.928,0 l-64.613,63.123l0.655-488.012c0-27.829-22.399-50.304-50.129-50.304c-27.734,0-50.129,22.474-50.129,50.304v598.873 c0,20.869,12.265,39.064,31.464,47.089c5.865,2.68,12.265,3.745,18.664,3.745c13.33,0,38.629-5.35,48.229-14.98L998.068,708.89 C1016.732,689.09,1016.732,656.986,997.533,637.187z"></path>
            </svg>
            <svg viewBox="0 0 1024 1024">
              <path d="M475.67,647.106L277.05,449.308c-20.067-19.968-20.067-52.394,0-72.277l0,0 c20.102-20.012,52.552-20.012,72.699-0.084l0,0l162.207,161.516l162.372-161.516c20.062-19.928,52.517-19.928,72.624,0l0,0 c20.062,19.968,20.062,52.394,0,72.362l0,0L548.334,647.106C538.76,656.504,525.573,662,511.957,662l0,0 C498.425,662,485.194,656.504,475.67,647.106L475.67,647.106z"></path>
            </svg>
          </div>
        </div>
      </div>
      <div class="main">
        <div class="items">

          <router-link :to="'/'+id + '/'+item.id" v-for="(item,index)  in  items" :key='index' :class="[{'active': item.id === listId}, 'item']">
            <div class="title">
              <div>{{ item.title }}</div>
              <div v-if="item.type !== 'note'" class="file-type">
                          <span>
                            <svg v-if="item.type === 'word'" viewBox="0 0 1024 1024"><path d="M535.119473 0h69.599248v95.247413C729.226717 96.331138 853.614299 93.92286 977.881468 96.331138a40.459078 40.459078 0 0 1 44.914393 45.516463c2.047037 234.566322 0 469.614299 1.204139 703.819379-1.204139 24.082785 2.287865 50.694262-11.318909 72.248354-16.978363 12.041392-38.893697 10.837253-58.761994 12.041392h-349.200376V1023.518344h-72.248354C354.980245 990.886171 177.490122 960.541863 0 928.752587V95.488241C178.33302 63.578551 356.786453 32.511759 535.119473 0z" fill="#2A5699" p-id="5505"></path><path d="M604.718721 131.010348H988.598307v761.979304H604.718721v-95.247413h302.479774v-48.165569H604.718721v-59.002822h302.479774v-48.16557H604.718721v-59.002822h302.479774v-48.165569H604.718721v-60.206961h302.479774V428.673565H604.718721v-60.206961h302.479774v-46.96143H604.718721v-59.604892h302.479774V214.336783H604.718721zM240.827846 341.373471c22.156162-1.324553 44.19191-2.287865 66.348071-3.492003 15.533396 80.4365 31.30762 160.632173 48.165569 240.827845 13.125118-82.724365 27.695202-165.087488 41.783632-247.571025 23.239887-0.842897 46.479774-2.167451 69.719661-3.612418-26.370649 115.356538-49.369708 231.796802-78.148636 346.430856-19.386642 10.355597-48.165569 0-71.52587 1.204139C301.034807 596.169332 283.093133 517.779868 269.245532 438.667921c-13.606773 76.944497-31.30762 153.16651-46.841016 229.508937-22.39699-1.204139-44.793979-2.528692-67.311383-4.094073-19.266228-104.760113-42.024459-208.918156-60.206962-313.919097 19.868297-0.963311 39.857008-1.806209 60.206962-2.528693 12.041392 75.860771 25.648166 151.360301 36.124177 227.341487 16.135466-77.907808 32.873001-155.695202 49.610536-233.603011z" fill="#FFFFFF" p-id="5506"></path></svg>
                            <svg v-if="item.type === 'pdf'" viewBox="0 0 1024 1024"><path d="M910.2336 1024H113.7664C51.2 1024 0 972.8 0 910.2336V113.7664C0 51.2 51.2 0 113.7664 0h796.4672C972.8 0 1024 51.2 1024 113.7664v796.4672C1024 972.8 972.8 1024 910.2336 1024zM155.2128 555.4432h44.672c35.9936 1.2032 66.0224-8.704 89.984-29.6704 24.0128-20.9664 36.0192-48.3328 36.0192-82.1248 0-33.28-10.24-58.9568-30.7712-76.8768-20.5312-17.92-49.3568-26.88-86.5536-26.88h-97.792v344.2432h44.4416v-128.6912z m0-176.4352h45.3376c52.48 0 78.6688 22.1184 78.6688 66.432 0 22.8096-7.04 40.3456-21.12 52.5824-14.08 12.1856-34.6624 18.304-61.7728 18.304H155.2128v-137.3184z m234.0352 305.1008h94.6432c54.528 0 99.072-16.0768 133.6576-48.2048 34.6112-32.1536 51.8912-74.9056 51.8912-128.256 0-51.0976-17.28-91.8528-51.8912-122.2144-34.56-30.336-77.952-45.568-130.0992-45.568h-98.2016v344.2432z m44.416-304.6656h51.9936c41.0624 0 74.1632 10.8032 99.328 32.4608 25.216 21.6064 37.7856 53.888 37.7856 96.8704 0 43.008-12.2368 76.3392-36.7872 100.224-24.5248 23.8592-58.624 35.7632-102.3232 35.7632H433.664V379.4432z m479.5648 0v-39.5776h-177.3312v344.2432h44.4416v-150.4512h123.1104V494.592h-123.136v-115.1232h132.9152z" fill="#D93838" p-id="13125"></path></svg>
                            <svg v-if="item.type === 'ppt'" viewBox="0 0 1024 1024"><path d="M968.704 135.168h-430.08v752.64h430.08c15.36 0 26.624-12.288 26.624-26.624V162.816c0-15.36-11.264-27.648-26.624-27.648z" fill="#FF8A65" p-id="11255"></path><path d="M512 646.144h376.832v53.248H512z m0 107.52h376.832v54.272H512z m161.792-483.328c-89.088 0-161.792 72.704-161.792 161.792C512 521.216 584.704 593.92 673.792 593.92s161.792-72.704 161.792-161.792H673.792V270.336z" fill="#FBE9E7" p-id="11256"></path><path d="M727.04 216.064v161.792h161.792c0-89.088-72.704-161.792-161.792-161.792z" fill="#FBE9E7" p-id="11257"></path><path d="M592.896 996.352L28.672 888.832V135.168L592.896 27.648z" fill="#E64A19" p-id="11258"></path><path d="M319.488 327.68H192.512v368.64h78.848V569.344h40.96c44.032 0 78.848-11.264 104.448-34.816 25.6-22.528 38.912-53.248 38.912-90.112C455.68 366.592 409.6 327.68 319.488 327.68z m-14.336 178.176h-33.792V391.168h33.792c43.008 0 64.512 18.432 64.512 56.32 0 39.936-21.504 58.368-64.512 58.368z" fill="#FFFFFF" p-id="11259"></path></svg>
                            <svg v-if="item.type === 'excel'" viewBox="0 0 1024 1024"><path d="M968.704 135.168h-430.08v752.64h430.08c15.36 0 26.624-12.288 26.624-26.624V162.816c0-15.36-11.264-27.648-26.624-27.648z" fill="#4CAF50" p-id="5257"></path><path d="M727.04 270.336h188.416v80.896H727.04z m0 268.288h188.416v80.896H727.04z m0 134.144h188.416v80.896H727.04z m0-268.288h188.416v80.896H727.04zM538.624 270.336h134.144v80.896H538.624z m0 268.288h134.144v80.896H538.624z m0 134.144h134.144v80.896H538.624z m0-268.288h134.144v80.896H538.624z" fill="#FFFFFF" p-id="5258"></path><path d="M592.896 996.352L28.672 888.832V135.168L592.896 27.648z" fill="#2E7D32" p-id="5259"></path><path d="M379.904 696.32l-63.488-119.808c-2.048-4.096-5.12-12.288-7.168-24.576H307.2c-1.024 6.144-4.096 14.336-8.192 25.6L235.52 696.32h-99.328l117.76-184.32-107.52-184.32h101.376l53.248 110.592c4.096 9.216 8.192 19.456 11.264 30.72h1.024c2.048-7.168 6.144-17.408 11.264-31.744L381.952 327.68h92.16L364.544 509.952 478.208 696.32h-98.304z" fill="#FFFFFF" p-id="5260"></path></svg>
                          </span>
              </div>
            </div>
            <div class="describe">
              <div class="info">
                <div class="tag">
                  <p class="time">2020/02/25</p>
                  <p>我的笔记</p>
                  <p>我的笔记</p>
                  <p>我的笔记</p>
                  <p>我的笔记</p>
                  <p>我的笔记</p>
                  <p>我的笔记</p>
                  <p>我的笔记</p>
                </div>
                <div v-if="item.type === 'note'" class="content">
                  <p>
                  </p>
                </div>
              </div>
              <div v-if="item.type === 'note'" class="img">
                <img src="http://walter.copyf.com/assets/img/faces/avatar.jpg" alt="">
              </div>
            </div>
            <div class="edit evocationModal">
              <svg viewBox="0 0 24 24">
                <path d="M13.5,17c0,0.8-0.7,1.5-1.5,1.5s-1.5-0.7-1.5-1.5s0.7-1.5,1.5-1.5S13.5,16.2,13.5,17z M13.5,12c0,0.8-0.7,1.5-1.5,1.5 s-1.5-0.7-1.5-1.5s0.7-1.5,1.5-1.5S13.5,11.2,13.5,12z M13.5,7c0,0.8-0.7,1.5-1.5,1.5S10.5,7.8,10.5,7s0.7-1.5,1.5-1.5 S13.5,6.2,13.5,7z"></path>
              </svg>
            </div>
          </router-link>

        </div>
      </div>
    </div>
    <router-view></router-view>
  </div>
</template>

<script>
    import {request} from '@/request/http'

    export default {
        name: 'lists',
        data() {
            return {
                id: '',
              listId: '',
                items: [],
            }
        },

      watch: {
        '$route.params.id'() {
          // 监听$route.params.id的变化，如果这个id即代表用户点击了其他的待办项需要重新请求数据。
          this.listId = this.$route.params.uuid;
          this.getLists()
        }
      },
        created() {
            this.getLists();
        },
        methods: {
            // 获取菜单数据
            getLists() {
              const ID = this.$route.params.id;
              this.id = ID;
                request({
                    url: "/",
                    data: 'query{list(cid:"' + ID + '") {id title type}}',
                }).then(res => {
                    this.items = res.data.list;
                }).catch(res => {

                })
            },

        }
    }
</script>

<style>
  #listContent {
    display: flex;
    flex: 1;
  }
    #list .header h1 {
        font-size: 2rem;
        color: rgb(51, 51, 51);
        font-weight: 400;
    }

    #list {
        width: 320px;
        height: 100vh;
        display: flex;
        flex-direction: column;
        background: rgb(245, 247, 249);
        border-right: 1px solid rgb(230, 236, 241);
        box-sizing: border-box;
    }

    #list .header {
        height: 80px;
        display: flex;
        align-items: center;
        justify-content: center;
        border-bottom: 1px solid #e6e9ed;
        box-sizing: border-box;
        position: relative;
    }

    #list .header .edit {
        position: absolute;
        top: 0;
        right: 0;
        padding: 10px;
        display: none;
    }

    #list .header:hover .edit {
        display: block;
    }

    #list .filter {
        height: 40px;
        display: flex;
        align-items: center;
        justify-content: center;
        border-bottom: 1px solid #e6e9ed;
        box-sizing: border-box;
    }

    #list .filter .main {
        display: flex;
        width: 87%;
    }

    #list .filter .selector-number {
        display: flex;
        box-sizing: border-box;
        border-radius: 3px;
        flex-grow: 1;
        align-items: center;
    }

    #list .filter .selector-number p {
        color: #172b4d;
        font-size: 14px;
        font-weight: 400;
        line-height: 1.6;
    }

    #list .filter .sort {
        display: flex;
        padding: 4px 8px;
        border-radius: 4px;
        color: rgb(80, 95, 121);
    }

    #list .filter .sort:hover {
        background-color: #e6e9ed;
    }

    #list > .main {
        flex-grow: 1;
        position: relative;
    }

    #list .main .items {
        position: absolute;
        top: 0px;
        left: 0px;
        right: 0px;
        bottom: 0px;
        overflow: scroll;
        scrollbar-width: none; /* firefox */
        -ms-overflow-style: none; /* IE 10+ */
        overflow-x: hidden;
        overflow-y: auto;
    }

    #list .main .items::-webkit-scrollbar {
        display: none; /* Chrome Safari */
    }

    #list .main .items .item {
        padding: 20px;
        border-bottom: 1px solid;
        border-bottom-color: #e6e9ed;
        box-sizing: border-box;
        display: flex;
        flex-direction: column;
        position: relative;
    }

    #list .main .items .item.router-link-active {
        background-color: #ffffff;
    }

    #list .main .items .item .edit {
        position: absolute;
        top: 0;
        right: 0;
        padding: 2px 0;
        display: none;
    }

    #list .main .items .item:hover .edit {
        display: block;
    }

    #list .main .items .item .title {
        height: 22px;
        display: flex;
        flex-wrap: nowrap;
        align-items: center;
        color: #172b4d;
        font-size: 14px;
        font-weight: 400;
        line-height: 1.6;
    }

    #list .main .items .item .title div {
        flex-grow: 1;
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
        display: flex;
    }


    #list .main .items .item .title .file-type {
        flex-grow: unset;
    }

    #list .main .items .item .title .file-type span {
        display: grid;
    }

    #list .main .items .item .describe {
        display: flex;
        align-items: center;
        flex-grow: 1;
        color: #a7afbc;
        font-size: 12px;
        font-weight: 400;
        line-height: 1.6;
    }

    #list .main .items .item .describe .info {
        display: flex;
        box-sizing: border-box;
        border-radius: 3px;
        flex-grow: 1;
        flex-direction: column;
    }

    #list .main .items .item .describe .info .tag {
        display: flex;
        padding-top: 4px;
    }

    #list .main .items .item .describe .info .tag p {
        flex-grow: 1;
        width: 10px;
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
    }

    #list .main .items .item .describe .info .tag .time {
        min-width: 67px;
        max-width: 67px;
        flex-shrink: 0;
        margin-right: 16px;
    }

    #list .main .items .item .describe .info .content {
        padding-top: 4px;
        display: -webkit-box;
        overflow: hidden;
        flex-grow: 1;
        max-height: 42px;
        text-overflow: ellipsis;
        -webkit-box-orient: vertical;
        -webkit-line-clamp: 2;
    }

    #list .main .items .item .describe .img {
        display: flex;
    }

    #list .main .items .item .describe .img img {
        width: 56px;
        height: 56px;
        margin-left: 12px;
    }
</style>