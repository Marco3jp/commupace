import Vue from 'vue';
import Router from 'vue-router';

import Community from '@view/community'
import CommunityFallback from '@view/communityFallBack'
import CommunityFrame from '@view/frame/communityFrame.vue'
import CommunityFront from '@view/communityFront.vue'
import CommunityChat from '@view/communityChat'
import CommunityThreadFrame from '@view/frame/communityThreadFrame'
import CommunityThreadFront from '@view/communityThreadFront'
import CommunityThread from '@view/communityThread'
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
                    component: CommunityFrame,
                    children: [
                        {
                            path: '',
                            name: 'communityFront',
                            component: CommunityFront,
                        },
                        {
                            path: 'chat',
                            name: 'communityChat',
                            component: CommunityChat,
                        },
                        {
                            path: 'thread',
                            component: CommunityThreadFrame,
                            children: [
                                {
                                    path: '',
                                    name: 'communityThreadFront',
                                    component: CommunityThreadFront,
                                },
                                {
                                    path: ':threadId',
                                    name: 'communityThread',
                                    component: CommunityThread,
                                }
                            ]
                        }
                    ]
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