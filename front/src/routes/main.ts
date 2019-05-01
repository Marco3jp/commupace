import Vue from 'vue';
import Router from 'vue-router';

import Community from '@view/community'
import Explore from '@view/explore'
import MyPage from '@view/mypage'

Vue.use(Router);

/**
 * path: "urlパス"
 * component: "使用しているコンポーネント"
 */

export default new Router({
    mode: 'history',
    routes: [
        {
            path: '/myPage',
            name: 'myPage',
            component: MyPage
        },
        {
            path: '/community',
            name: 'community',
            component: Community
        },
        {
            path: '/explore',
            name: 'explore',
            component: Explore
        },
    ]
});