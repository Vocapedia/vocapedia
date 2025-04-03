import languages from './languages.json';

export function getLangByCode(code) {
    return languages.find(lang => lang.code === code) || null;
}
