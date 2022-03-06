type GoogleAnalyticsContactEvent = {
  action: 'submit_form';
  category: 'contact';
  label: string;
};

type GoogleAnalyticsClickEvent = {
  action: 'click';
  category: 'other';
  label: string;
};

type GoogleAnalyticsEvent = (GoogleAnalyticsContactEvent | GoogleAnalyticsClickEvent) & {
  label?: Record<string, string | number | boolean>;
  value?: string;
};

export type { GoogleAnalyticsEvent };
