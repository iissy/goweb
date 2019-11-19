<template>
  <div style="position: relative;">
    <Menu tagIndex="4"></Menu>
    <div class="rightMain">
      <div style="padding:0 0 0 0;height:60px;margin-bottom:20px;">
        <div style="background-color: #ffffff;height:60px;padding:10px;"></div>
      </div>
      <div id="list">
        <form class="form-horizontal mg0">
          <div id="dataform" class="row">
            <div class="col-md-12">
              <div class="portlet light" style="margin-bottom: 0;">
                <div class="portlet-body form" style="margin:0 auto;">
                  <div class="form-horizontal mg0" role="form">
                    <div class="form-body" style="padding:0 0 0 0;">
                      <div class="form-group">
                        <label class="col-md-1 control-label">标题</label>
                        <div class="col-md-7">
                          <input
                            type="text"
                            v-model="Subject"
                            class="form-control"
                            placeholder="请在此填写标题"
                          />
                        </div>
                        <div class="col-md-2">
                          <select v-model="PostType" class="form-control">
                            <option value="技术">技术</option>
                            <option value="科技">科技</option>
                            <option value="新闻">新闻</option>
                            <option value="故事">故事</option>
                          </select>
                        </div>
                        <div class="col-md-2">
                          <select v-model="Origin" class="form-control">
                            <option value="爱施缘">爱施缘</option>
                            <option value="中关村在线">中关村在线</option>
                            <option value="凤凰网科技">凤凰网科技</option>
                            <option value="新浪科技">新浪科技</option>
                            <option value="网易科技">网易科技</option>
                            <option value="快科技">快科技</option>
                            <option value="搜狐科技">搜狐科技</option>
                            <option value="太平洋电脑网">太平洋电脑网</option>
                          </select>
                        </div>
                      </div>
                      <div class="form-group">
                        <label class="col-md-1 control-label">连接</label>
                        <div class="col-md-11">
                          <input
                            type="text"
                            id="Id"
                            v-model="Id"
                            class="form-control"
                            placeholder="请在此填写url链接"
                          />
                        </div>
                      </div>
                      <div class="form-group">
                        <label class="col-md-1 control-label">图片</label>
                        <div class="col-md-11">
                          <div
                            id="imgfileinput"
                            data-provides="fileinput"
                            class="fileinput fileinput-new right"
                          >
                            <span class="btn green btn-file">
                              <span class="fileinput-new"> 选择图片 </span>
                              <input
                                type="file"
                                id="file"
                                @change="change"
                                name="image"
                              />
                            </span>
                          </div>
                          <input
                            type="text"
                            v-model="Picture"
                            class="form-control"
                          />
                        </div>
                      </div>
                      <div class="form-group">
                        <label class="col-md-1 control-label">简介</label>
                        <div class="col-md-11">
                          <textarea
                            class="form-control"
                            v-model="Description"
                            rows="2"
                          ></textarea>
                        </div>
                      </div>
                      <div class="form-group">
                        <label class="col-md-1 control-label">文章内容</label>
                        <div class="col-md-11">
                          <vue-html5-editor
                            :content="Body"
                            :height="508"
                            :auto-height="false"
                            :show-module-name="showModuleName"
                            @change="updateData"
                            ref="editor"
                          ></vue-html5-editor>
                        </div>
                      </div>
                      <div class="form-group">
                        <label class="col-md-1 control-label"></label>
                        <div class="col-md-11">
                          <button
                            style="color:#ffffff;padding:5px 50px 5px 50px;background-color: #36c6d3;font-size:16px;border-radius:10px;border: 1px solid #2bb8c4;float:right;"
                            type="button"
                            v-on:click="post"
                          >
                            发 布
                          </button>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import Menu from "../components/menu";
import Vue from "vue";
import VueHtml5Editor from "vue-html5-editor";
import router from "../../router";
import httper from "../../util/httper";
import "../../util/upload";

Vue.use(VueHtml5Editor, {
  showModuleName: true,
  image: {
    sizeLimit: 512 * 1024,
    upload: {
      url: "/upload"
    }
  }
});

export default {
  data: function() {
    return {
      Id: "",
      PostType: "",
      Origin: "",
      Subject: "",
      Picture: "",
      Description: "",
      Body: "",
      Adding: true,
      showModuleName: false
    };
  },
  components: {
    Menu
  },
  created: function() {
    var self = this;
    if (self.$route.params.id) {
      var url = "/article/get/" + self.$route.params.id;
      httper
        .get(url)
        .then(function(response) {
          self.Id = response.data.ID;
          self.PostType = response.data.PostType;
          (self.Origin = response.data.Origin),
            (self.Subject = response.data.Subject);
          self.Picture = response.data.Picture;
          self.Description = response.data.Description;
          self.Body = response.data.Body;
          self.Adding = false;
          $("#Id").attr("disabled", true);
        })
        .catch(function(error) {
          console.log(error);
        });
    }
  },
  methods: {
    updateData: function(data) {
      this.Body = data;
    },
    fullScreen: function() {
      this.$refs.editor.enableFullScreen();
    },
    focus: function() {
      this.$refs.editor.focus();
    },
    change: function() {
      var self = this;
      $("#file").upload("/upload", function(response) {
        self.Picture = response.data;
      });
    },
    post: function() {
      var self = this;
      if (
        !!$.trim(self.Picture) &&
        !!$.trim(self.Id) &&
        !!$.trim(self.Subject) &&
        !!$.trim(self.Body) &&
        !!$.trim(self.Description)
      ) {
        httper
          .post("/article/post", {
            Id: self.Id,
            PostType: self.PostType,
            Origin: self.Origin,
            Subject: self.Subject,
            Picture: self.Picture,
            Description: self.Description,
            Body: self.Body,
            Adding: self.Adding
          })
          .then(function(response) {
            if (response.data.ok) {
              router.push({
                name: "ArticleList",
                params: { size: 15, pageno: 1 }
              });
            }
          })
          .catch(function(error) {
            console.log(error);
          });
      }
    }
  }
};
</script>