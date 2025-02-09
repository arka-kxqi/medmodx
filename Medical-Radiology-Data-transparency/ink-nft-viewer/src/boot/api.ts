import { capitalize, objToArray } from '@astar-network/astar-sdk-core';
import { ApiPromise } from '@polkadot/api';
import { keyring } from '@polkadot/ui-keyring';
import { useMeta } from 'quasar';
import { boot } from 'quasar/wrappers';
import { connectApi } from 'src/config/api/polkadot/connectApi';
import { endpointKey, providerEndpoints } from 'src/config/chainEndpoints';
import { LOCAL_STORAGE } from 'src/config/localStorage';
import { opengraphMeta } from 'src/config/metadata';
import { isMobileDevice } from 'src/modules/wallet';
import { useChainInfo } from 'src/hooks/useChainInfo';
import { useExtensions } from 'src/hooks/useExtensions';
import { useMetaExtensions } from 'src/hooks/useMetaExtensions';
import { SubstrateAccount } from 'src/store/general/state';
import { computed, ref, watchPostEffect } from 'vue';

let $api: ApiPromise | undefined;
const $endpoint = ref<string>('');

export default boot(async ({ store }) => {
  const { NETWORK_IDX, CUSTOM_ENDPOINT, SELECTED_ENDPOINT, SELECTED_ADDRESS } = LOCAL_STORAGE;

  const networkIdxStore = localStorage.getItem(NETWORK_IDX);
  const customEndpoint = localStorage.getItem(CUSTOM_ENDPOINT);
  const selectedEndpointData = localStorage.getItem(SELECTED_ENDPOINT);
  if (!selectedEndpointData) {
    if (networkIdxStore !== null) {
      const networkIdx = Number(networkIdxStore);
      localStorage.setItem(
        LOCAL_STORAGE.SELECTED_ENDPOINT,
        JSON.stringify({ [networkIdx]: providerEndpoints[networkIdx].endpoints[0].endpoint })
      );
    } else {
      localStorage.setItem(
        LOCAL_STORAGE.SELECTED_ENDPOINT,
        JSON.stringify({ '0': providerEndpoints[0].endpoints[0].endpoint })
      );
    }
  }
  const selectedAddress = localStorage.getItem(SELECTED_ADDRESS);
  const selectedEndpoint = selectedEndpointData ? JSON.parse(selectedEndpointData) : {};
  if (networkIdxStore) {
    store.commit('general/setCurrentNetworkIdx', Number(networkIdxStore));
  }
  if (customEndpoint) {
    store.commit('general/setCurrentCustomEndpoint', customEndpoint);
  }
  const networkIdx = computed(() => store.getters['general/networkIdx']);
  const defaultEndpoint = providerEndpoints[networkIdx.value].endpoints[0].endpoint;
  let endpoint = selectedEndpoint.hasOwnProperty(networkIdx.value)
    ? selectedEndpoint[networkIdx.value]
    : defaultEndpoint;
  if (networkIdx.value === endpointKey.CUSTOM) {
    const customEndpoint = computed(() => store.getters['general/customEndpoint']);
    endpoint = customEndpoint.value;
  }

  if (networkIdx.value === endpointKey.LOCAL) {
    endpoint = providerEndpoints[networkIdx.value].endpoints[0].endpoint;
  }

  // set metadata header
  const favicon = providerEndpoints[Number(networkIdx.value)].defaultLogo;
  const displayName = providerEndpoints[Number(networkIdx.value)].displayName;
  const networkName = capitalize(providerEndpoints[Number(networkIdx.value)].networkAlias);
  useMeta({
    title: '',
    // titleTemplate: (title) => `${title} | ${networkName} Portal - ${displayName}`,
    htmlAttr: { lang: 'en' },
    link: {
      material: {
        rel: 'icon',
        href: favicon,
      },
    },
    meta: opengraphMeta(displayName, networkName),
  });
  let { api } = await connectApi(endpoint, networkIdx.value, store);
  $api = api;
  $endpoint.value = endpoint;

  const seen = new Set();
  const accountMap: SubstrateAccount[] = [];
  keyring.accounts.subject.subscribe((accounts) => {
    if (accounts) {
      const accountArray = objToArray(accounts);
      accountArray.forEach((account) => {
        const { address, meta } = account.json;
        const source = meta.source;
        const addressWithSource = address + source;
        const isSeen = seen.has(addressWithSource);

        if (!isSeen) {
          const data = {
            address,
            name: meta.name.replace('\n              ', ''),
            source,
          };

          seen.add(addressWithSource);
          accountMap.push(data);
          store.commit('general/setSubstrateAccounts', accountMap);
        }
      });
    }
  });

  // update chaininfo
  const { chainInfo } = useChainInfo(api);
  watchPostEffect(async () => {
    store.commit('general/setChainInfo', chainInfo.value);
  });

  // execute extension process automatically if selectedAddress is linked or mobile device
  if (selectedAddress !== null || isMobileDevice) {
    const { extensions } = useExtensions(api, store);
    const { metaExtensions, extensionCount } = useMetaExtensions(api, extensions)!!;
    watchPostEffect(async () => {
      store.commit('general/setMetaExtensions', metaExtensions.value);
      store.commit('general/setExtensionCount', extensionCount.value);
    });
  }
});

export { $api, $endpoint };
