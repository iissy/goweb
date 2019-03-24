<template>
    <div class="clearfix">
        <div class="datainfo fl">第 <span id="PageNumber">{{current}}</span> 页 ( 总共 <span id="PageCount">{{page}}</span> 页 )</div>
        <nav class="fr">
            <ul class="pagination">
                <li :class="{'disabled': current == 1}"><a href="javascript:;" @click="setCurrent(1)">&laquo;</a></li>
                <li :class="{'disabled': current == 1}"><a href="javascript:;" @click="setCurrent(current - 1)">&lsaquo;</a></li>
                <li v-for="p in grouplist" :key="p.val" :class="{'active': current == p.val}">
                    <a href="javascript:;" @click="setCurrent(p.val)">{{ p.text }}</a>
                </li>
                <li :class="{'disabled': current == page}"><a href="javascript:;" @click="setCurrent(current + 1)">&rsaquo;</a></li>
                <li :class="{'disabled': current == page}"><a href="javascript:;" @click="setCurrent(page)">&raquo;</a></li>
            </ul>
        </nav>
    </div>
</template>

<script>
    export default {
        data() {
            return {
            };
        },
        props: {
            total: {// 数据总条数
                type: Number,
                default: 0
            },
            display: {// 每页显示条数
                type: Number,
                default: 10
            },
            current: {// 当前页码
                type: Number,
                default: 1
            },
            pagegroup: {// 分页条数
                type: Number,
                default: 5,
                coerce: function (v) {
                    v = v > 0 ? v : 5;
                    return v % 2 === 1 ? v : v + 1;
                }
            }
        },
        computed: {
            page: function () { // 总页数
                return Math.ceil(this.total / this.display);
            },
            grouplist: function () { // 获取分页页码
                var len = this.page, temp = [], list = [], count = Math.floor(this.pagegroup / 2), center = this.current;
                if (len <= this.pagegroup) {
                    while (len--) {
                        temp.push({ text: this.page - len, val: this.page - len });
                    }
                    return temp;
                }
                while (len--) {
                    temp.push(this.page - len);
                }

                var idx = temp.indexOf(center);
                (idx < count) && (center = center + count - idx);
                (this.current > this.page - count) && (center = this.page - count);
                temp = temp.splice(center - count - 1, this.pagegroup);
                do {
                    var t = temp.shift();
                    list.push({
                        text: t,
                        val: t
                    });
                } while (temp.length);
                if (this.page > this.pagegroup) {
                    (this.current > count + 1) && list.unshift({ text: '...', val: list[0].val - 1 });
                    (this.current < this.page - count) && list.push({ text: '...', val: list[list.length - 1].val + 1 });
                }
                return list;
            }
        },
        methods: {
            setCurrent: function (idx) {
                if (this.current !== idx && idx > 0 && idx < this.page + 1) {
                    this.current = idx;
                    this.$emit('pagechange', this.current);
                }
            }
        }
    };
</script>