<template>
    <div id="nav">
        <div class="header">
            <div class="item"  @click="newNote">
                <svg viewBox="0 0 1024 1024">
                    <path d="M812,461.999H562v-250c0-27.615-22.385-50-50-50s-50,22.385-50,50v250H212c-27.615,0-50,22.385-50,50
 c0,27.616,22.385,50.001,50,50.001h250v250c0,27.615,22.385,50,50,50s50-22.385,50-50v-250h250c27.615,0,50-22.385,50-50.001
 C862,484.384,839.615,461.999,812,461.999z"></path>
                </svg>
            </div>
            <div class="item">
                <div>
                    <el-button type="text" @click="open">
                        <svg viewBox="0 0 1024 1024">
                            <path d="M1015.456 273.056c-8.544-12.8-21.344-17.056-34.144-17.056h-196.256L627.2 98.144c-8.544-8.544-17.056-12.8-29.856-12.8H42.688C17.088 85.344 0.032 102.4 0.032 128v768c0 25.6 17.056 42.656 42.656 42.656h810.656c21.344 0 38.4-12.8 42.656-34.144l128-597.344c0-12.8 0-25.6-8.544-34.144z m-930.112-102.4h494.944L665.632 256H170.688c-21.344 0-38.4 12.8-42.656 34.144l-42.656 204.8V170.688zM819.2 853.344H93.856l110.944-512h725.344l-110.944 512z" p-id="2538"></path>
                        </svg>
                    </el-button>
                </div>
                <div>
                    <el-upload
                            class="item"
                            action="/uploadFile/"
                            :show-file-list="false"
                            :data="{id:menuId}"
                    >
                    <svg viewBox="0 0 1024 1024">
                        <path d="M662.63 545.37l-128-128a32 32 0 0 0-45.25 0l-128 128a32 32 0 1 0 45.25 45.25L480 517.25V872a32 32 0 0 0 64 0V517.25l73.37 73.37a32 32 0 1 0 45.25-45.25zM752 760H640a32 32 0 0 1 0-64h112c79.4 0 144-64.6 144-144a144 144 0 0 0-132.45-143.54 32 32 0 0 1-29.24-28C720.6 268.46 625 184 512 184s-208.6 84.46-222.31 196.45a32 32 0 0 1-29.24 28A144 144 0 0 0 128 552c0 79.4 64.6 144 144 144h112a32 32 0 0 1 0 64H272c-114.69 0-208-93.31-208-208a208.08 208.08 0 0 1 166.23-203.79C258 216.5 375 120 512 120s254 96.5 281.77 228.21A208.08 208.08 0 0 1 960 552c0 114.69-93.31 208-208 208z" p-id="3596"></path>
                    </svg>
                    </el-upload>
                </div>
            </div>
        </div>
        <div class="main">
            <div>
                <el-tooltip v-for="(item,index)  in  items" :key='index' effect="dark"
                            :content="item.title" placement="right">
                    <router-link :to="'/' + item.id" :class="[{'active': item.id === menuId}, 'item color-' +item.color]">
                        {{item.title.slice(0,1)}}
                    </router-link>
                </el-tooltip>
            </div>
        </div>
        <div class="footer">
            <el-tooltip effect="dark" content="近期笔记" placement="right">
            <router-link to="/all" :class="[{'active': 'all' === menuId},'item']">
                <svg viewBox="0 0 1024 1024">
                    <path d="M512.005,62C263.477,62,62,263.476,62,512c0,248.534,201.477,450,450.005,450
C760.528,962,962,760.534,962,512C962,263.476,760.528,62,512.005,62z M662.148,561.742H512.08c0,0-12.506,0-25.011,0
c-16.152,0-25.011-12.497-25.011-23.434v-226.5c0-15.681,7.333-28.318,24.941-28.647c16.843-0.315,25.071,14.171,25.081,28.647
v199.947h150.068c15.642,0,23.196,8.718,23.196,24.393C685.344,551.834,677.79,561.742,662.148,561.742z"></path>
                </svg>
            </router-link>
            </el-tooltip>
            <el-tooltip effect="dark" content="近期笔记" placement="right">
            <router-link to="/file" :class="[{'active': 'file' === menuId},'item']" title="文件">
                <span>
                    <svg viewBox="0 0 1024 1024">
                        <path d="M288.010553 735.989447h383.981325c17.672498 0 31.998785 14.326287 31.998784 31.998785s-14.326287 31.998785-31.998784 31.998785H288.010553c-17.672498 0-31.998785-14.326287-31.998785-31.998785s14.326287-31.998785 31.998785-31.998785z m0-159.992901h191.990662c17.672498 0 31.998785 14.326287 31.998785 31.998785s-14.326287 31.998785-31.998785 31.998785H288.010553c-17.672498 0-31.998785-14.326287-31.998785-31.998785s14.326287-31.998785 31.998785-31.998785zM636.111692 92.137467l231.753272 231.753271a95.994308 95.994308 0 0 1 28.116361 67.87897v472.212832c0 53.016471-42.97886 95.995331-95.995332 95.995331H224.014007c-53.016471 0-95.995331-42.97886-95.995332-95.995331V160.016437c0-53.016471 42.97886-95.995331 95.995332-95.995331h344.218715c25.459858 0 49.875944 10.114358 67.87897 28.116361z m-28.116361 62.388932v197.4807h197.4807L607.995331 154.526399z m-53.016471-26.507724H224.014007c-17.672498 0-31.998785 14.326287-31.998785 31.998785v703.966103c0 17.672498 14.326287 31.998785 31.998785 31.998785h575.971986c17.672498 0 31.998785-14.326287 31.998785-31.998785V405.024593v10.980076H607.995331c-35.344996 0-63.996546-28.652574-63.996546-63.996547V128.018675h10.980075z"></path>
                    </svg>
                </span>
            </router-link>
            </el-tooltip>
            <el-tooltip effect="dark" content="音视频" placement="right">

            <router-link to="/video" :class="[{'active': 'video' === menuId},'item']" title="音视频">
                <span>
                    <svg viewBox="0 0 1024 1024">
                        <path d="M960 192l-28.384 0c-16.8 0-32.928 6.624-44.928 18.432l-86.688 85.504 0-39.936c0-53.024-43.008-96-96-96l-608 0c-52.928 0-96 43.04-96 96l0 512c0 52.992 42.976 96 96 96l608 0c52.992 0 96-43.008 96-96l0-39.072 86.688 85.504c12 11.808 28.128 18.432 44.928 18.432l28.384 0c35.328 0 64-28.64 64-64l0-512.864c0-35.36-28.672-64-64-64zM96 800c-17.664 0-32-14.368-32-32l0-512c0-17.696 14.304-32 32-32l608 0c17.632 0 32 14.336 32 32l0 512c0 17.632-14.368 32-32 32l-608 0zM960 768.864l-32 0-128-128 0-0.864-32-32 0-192 160-160 32 0 0 512.864z"
                              p-id="2657"></path>                    </svg>
                </span>
            </router-link>
            </el-tooltip>
            <el-tooltip effect="dark" content="图片" placement="right">

            <router-link to="/image" :class="[{'active': 'image' === menuId},'item']" title="图片">
                <span>
                    <svg viewBox="0 0 1024 1024">
                        <path d="M959.87712 128c0.04096 0.04096 0.08192 0.08192 0.12288 0.12288l0 767.77472c-0.04096 0.04096-0.08192 0.08192-0.12288 0.12288l-895.77472 0c-0.04096-0.04096-0.08192-0.08192-0.12288-0.12288l0-767.77472c0.04096-0.04096 0.08192-0.08192 0.12288-0.12288l895.77472 0zM960 64l-896 0c-35.20512 0-64 28.79488-64 64l0 768c0 35.20512 28.79488 64 64 64l896 0c35.20512 0 64-28.79488 64-64l0-768c0-35.20512-28.79488-64-64-64l0 0zM832 288.01024c0 53.02272-42.98752 96.01024-96.01024 96.01024s-96.01024-42.98752-96.01024-96.01024 42.98752-96.01024 96.01024-96.01024 96.01024 42.98752 96.01024 96.01024zM896 832l-768 0 0-128 224.01024-384 256 320 64 0 224.01024-192z"
                              p-id="3376"></path>
                    </svg>
                </span>
            </router-link>
            </el-tooltip>
            <el-tooltip effect="dark" content="回收站" placement="right">
            <router-link to="/" class="item">
                <span>
                    <svg viewBox="0 0 1024 1024">
                        <path d="M186.996,912.005c0,26.575,22.725,49.995,50.005,49.995h550.008
 c25.995,0,49.985-23.42,49.985-49.995l0.005-650.004H187.001L186.996,912.005z M637.448,440.987
 c0-15.705,11.16-28.545,24.805-28.545c13.64,0,24.81,12.84,24.81,28.545v346.507c0,13.635-11.17,24.79-24.81,24.79
 c-13.645,0-24.805-11.155-24.805-24.79V440.987z M486.237,440.987c0-15.705,11.515-28.545,25.605-28.545
 c14.08,0,25.6,12.84,25.6,28.545v346.507c0,13.635-11.52,24.79-25.6,24.79c-14.09,0-25.605-11.155-25.605-24.79V440.987z
 M336.657,440.987c0-15.705,11.16-28.545,24.815-28.545c13.645,0,24.805,12.84,24.805,28.545v346.507
 c0,13.635-11.16,24.79-24.805,24.79c-13.655,0-24.815-11.155-24.815-24.79V440.987z M836.999,112H587.003
 c0-27.615-22.385-50-50-50h-50c-27.615,0-50,22.385-50,50H187.001c-27.615,0-50,22.385-50,50s22.385,50,50,50h649.998
 c27.615,0,50-22.385,50-50S864.614,112,836.999,112z"></path>
                    </svg>
                </span>
            </router-link>
            </el-tooltip>
        </div>
    </div>
</template>

<script>
    import {request} from '@/request/http'

    export default {
        name: 'menus',
        data() {
            return {
                menuId: 'all',
                type: 'note',
                items: [],
                model: ''
            }
        },

        created() {
            this.getMenus();
        },

        watch: {
            '$route.params.id'() {
                // 监听$route.params.id的变化，如果这个id即代表用户点击了其他的待办项需要重新请求数据。
                this.menuId = this.$route.params.id;
            }
        },
        methods: {
            // 获取菜单数据
            getMenus() {
                request({
                    url: "/",
                    data: 'query{list(cid:0) {id title color}}',
                }).then(res => {
                    this.items = res.data.list;
                }).catch(res => {

                })
            },
            open() {
                this.$prompt('', '文件夹名称', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    // inputPattern: /[\w!#$%&'*+/=?^_`{|}~-]+(?:\.[\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\w](?:[\w-]*[\w])?\.)+[\w](?:[\w-]*[\w])?/,
                    inputErrorMessage: '邮箱格式不正确'
                }).then(({ value }) => {
                    this.$message({
                        type: 'success',
                        message: '文件夹名称是: ' + value
                    });
                    request({
                        url: "/",
                        data: 'mutation{create(title:"'+ value +'") {id title}}',
                    }).then(res => {
                        this.getMenus();
                    }).catch(res => {

                    });

                }).catch(() => {
                    this.$message({
                        type: 'info',
                        message: '取消输入'
                    });
                });
            },
            newNote(){
                this.$prompt('', '笔记名称', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                }).then(({ value }) => {
                    this.$message({
                        type: 'success',
                        message: '笔记是: ' + value
                    });
                    request({
                        url: "/",
                        data: 'mutation{create(cid:'+ this.menuId +' title:"'+ value +'") {id title}}',
                    }).then(res => {
                        this.getMenus();
                    }).catch(res => {

                    });
                }).catch(() => {
                    this.$message({
                        type: 'info',
                        message: '取消输入'
                    });
                });
            }

        }
    }
</script>


<style scoped>


    #nav {
        width: 50px;
        height: 100vh;
        display: flex;
        flex-shrink: 0;
        flex-direction: column;
        color: rgb(157, 170, 182);
        background: rgb(255, 255, 255);
        border-right: 1px solid rgba(0, 0, 0, 0.12);
    }

    #nav .header {
        height: 80px;
        display: flex;
        flex-direction: column;
        justify-content: center;
        border-bottom: 1px solid #e6e9ed;
        box-sizing: border-box;
    }

    #nav .header .item {
        height: 25px;
        background-color: #448aff;
        color: #ffffff;
    }

    #nav .header .item>div {
        flex-grow: 1;
        height: 25px;
        display: -webkit-box;
        display: -ms-flexbox;
        display: flex;
        -webkit-box-align: center;
        -ms-flex-align: center;
        align-items: center;
        -webkit-box-pack: center;
        -ms-flex-pack: center;
        justify-content: center;
    }

    #nav .header .item svg {
        color: #ffffff;
    }

    #nav .header .item:first-child:hover {
        background-color: rgba(68, 138, 255, 0.8);
    }

    #nav .item {
        height: 40px;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    #nav .item.router-link-active, #nav .main .item:hover, #nav .footer .item:hover {
        background-color: rgba(90, 95, 121, 0.08);
    }

    #nav .main {
        flex-grow: 1;
        position: relative;
    }

    #nav .main > div {
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

    #nav .main > div::-webkit-scrollbar {
        display: none; /* Chrome Safari */
    }

    #nav .footer {
        min-height: 40px;
        border-top: 1px solid #e6e9ed;
        box-sizing: border-box;
    }
</style>