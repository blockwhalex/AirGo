import { createApp } from 'vue';
import pinia from '/@/stores/index';
import App from '/@/App.vue';
import router from '/@/router';
import { directive } from '/@/directive/index';
import other from '/@/utils/other';

import ElementPlus from 'element-plus';
import * as ElementPlusIconsVue from '@element-plus/icons-vue' //Element Plus 常用的图标集合
import '/@/theme/index.scss'; //样式


const app = createApp(App);
directive(app);
other.elSvg(app);

for (const [key, component] of Object.entries(ElementPlusIconsVue)) { //Element Plus 常用的图标集合
    app.component(key, component)
}
app.use(pinia).use(router).use(ElementPlus).mount('#app');
