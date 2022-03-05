import en from '../locales/en';
import ja from '../locales/ja';
import { useRouter } from 'next/router';

export const useLocale = () => {
  const { locale } = useRouter();
  const t = locale === 'ja' ? ja : en;
  return { locale, t };
};
