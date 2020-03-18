<template>
  <div id="content">
    <div id="editor">
      <froala v-if="type == 'note'" id="edit" :config="config" v-model="model"></froala>
      <iframe v-if="type == 'file'" :src="like"></iframe>
      <div v-if="type == 'image'">

      </div>
    </div>
  </div>
</template>


<script>
  import {request} from '@/request/http'

  export default {
    name: 'editor',
    data() {
      return {
        type:'note',
        like:'',
        model:'',
        config: {
          height: (document.documentElement.clientHeight - 90),
          language: 'zh_cn',
          toolbarInline: false,
          enter: false,
          imageButtons: ["floatImageLeft", "floatImageNone", "floatImageRight", "linkImage", "replaceImage", "removeImage"],

          requestHeaders: {
            // 'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
          },
          imageUploadURL: '/uploadFile/',
          videoUploadURL: '/',
          placeholderText: null,
          imageManagerLoadURL: "",
          imageManagerDeleteURL: "",
          saveInterval: 100,

          // Set the save param.
          // saveParam: 'content',

          // Set the save URL.
          // saveURL: './json.json',

          // HTTP request type.
          // saveMethod: 'POST',

          // Additional save params.
          // saveParams: {id: 'my_editor'},

          events: {
            'contentChanged': function () {
              // alert(this)
            },
          }
        },
      }
    },

    watch: {
      '$route.params.uuid'() {
        // 监听$route.params.id的变化，如果这个id即代表用户点击了其他的待办项需要重新请求数据。
        this.getEditor()
      }
    },
    created() {
      this.getEditor();
    },
    methods: {
      // 获取菜单数据
      getEditor() {
        const ID = this.$route.params.uuid;
        request({
          url: "/",
          data: 'query{note(id:"' + ID + '") {id title type like}}',
        }).then(res => {
          this.type = res.data.note.type;
          this.like = res.data.note.like;
        }).catch(res => {

        })
      }

    }
  }
</script>
<style>
  #content {
    display: flex;
    height: 100vh;
    padding: 0;
    flex-grow: 1;
    min-width: 0;
    flex-direction: column;
    background: rgb(255, 255, 255);
    flex-flow: column;
  }
  #content .header {
    height: 80px;
    display: flex;
    flex-flow: column;
    align-items: center;
    justify-content: center;
    border-bottom: 1px solid #e6e9ed;
    box-sizing: border-box;
  }

  #content .header .main {
    display: flex;
    width: 87%;
  }

  #content .header .title {
    display: flex;
    box-sizing: border-box;
    border-radius: 3px;
    flex-grow: 1;
    align-items: center;
    color: #172b4d;
    font-size: 20px;
    font-weight: 400;
    line-height: 1.3;
  }

  #content .header .buttons {
    display: flex;
    padding: 4px 8px;
    border-radius: 4px;
  }

  #content .header .buttons span {
    padding: 8px;
  }

  #content #editor {
    flex-grow: 1;
  }

  a[href="https://froala.com/wysiwyg-editor"], a[href="https://www.froala.com/wysiwyg-editor?k=u"] {
    display: none !important;
    position: absolute;
    top: -99999999px;
  }

  #editor {
    width: 100%;
    height: 100%;
  }

  #editor iframe {
    width: 100%;
    height: 100%;
    border: none;
  }

  #editor .fr-wrapper {
    border: none
  }

  #editor .second-toolbar {
    border: none;
  }

  #editor .second-toolbar #logo {
    display: none;
  }

  #editor .fr-toolbar.fr-desktop.fr-top.fr-basic {
    border: none
  }

  #editor .fr-toolbar.fr-desktop.fr-top.fr-basic.fr-sticky-on {
    top: 80px !important;
  }

  #editor .fr-toolbar.fr-desktop.fr-top.fr-basic.fr-sticky-off {
    top: 0px !important;
  }

  #editor .fr-toolbar .fr-newline {
    margin: 0;
  }

  .fr-box.fr-basic .fr-element {
    padding-bottom: 0;
  }
</style>