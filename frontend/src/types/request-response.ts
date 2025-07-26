export type KeyValue = {
    key: string;
    value: string;
  };
  
  export type GoResponse = {
    body: string;
    headers: KeyValue[];
    "status-code": number;
  };
  