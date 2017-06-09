// @flow
import path from 'path'

const commands = {
  'local-visdiff': {
    env: {
      KEYBASE_JS_VENDOR_DIR: process.env['KEYBASE_JS_VENDOR_DIR'] || path.resolve('../../js-vendor-desktop'),
      VISDIFF_DRY_RUN: 1,
    },
    help: 'Perform a local visdiff',
    shell: 'cd ../visdiff && yarn install --pure-lockfile && cd ../shared && node ../visdiff/dist/index.js',
  },
  'render-screenshots': {
    env: {
      BABEL_ENV: 'yarn',
      ELECTRON_ENABLE_LOGGING: 1,
      KEYBASE_NO_ENGINE: 1,
      VISDIFF: 'true',
    },
    help: 'Render images of dumb components',
    shell: 'yarn run _helper build-dev && electron ./desktop/dist/render-visdiff.bundle.js',
  },
}

export default commands
