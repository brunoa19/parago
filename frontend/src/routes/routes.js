import DesignPage from '../pages/DesignPage'
import LoginPage from '../pages/Auth/Login'
import { ROUTER_PATHS } from '../utils/constants'
import DetailsPage from '../pages/DetailsPage'
import SearchPage from '../pages/SearchPage'

const routes = Object.entries({
  designPage: {
    path: ROUTER_PATHS.designPage,
    component: DesignPage,
    exact: true,
  },
  searchPage: {
    path: ROUTER_PATHS.searchPage,
    component: SearchPage,
    exact: true,
  },
  detailsPage: {
    path: `${ROUTER_PATHS.detailsPage}/:configId`,
    component: DetailsPage,
    exact: true,
  },
  loginPage: {
    path: ROUTER_PATHS.loginPage,
    component: LoginPage,
    exact: true,
  },
})

export default routes
