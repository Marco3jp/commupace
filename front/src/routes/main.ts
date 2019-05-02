import Vue from 'vue';
import Router from 'vue-router';

import Community from '@view/community'
import CommunityFallback from '@view/communityFallBack'
import CommunityFront from '@view/communityFront.vue'
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
            component: Community,
            children: [
                {
                    path: '',
                    component: CommunityFallback,
                    name: 'community',
                },
                {
                    path: ':communityId',
                    name: 'communityFront',
                    component: CommunityFront
                }
            ]
        },
        {
            path: '/explore',
            name: 'explore',
            component: Explore
        },
    ]
});