import { getElByKey } from 'src/helpers/alerts';
import { describe, it, expect, vi, afterEach } from 'vitest';

describe('helpers/alerts', () => {
  
  describe('getElByKey', () => {
    const getElByKeyWithSpy = vi.fn().mockImplementation(getElByKey)

    afterEach(() => {
      getElByKeyWithSpy.mockReset();
    });

    it('should return the first key passed and only be called once when passed an array with a single string element', () => {
      const schemaKeys = ['error_name'];
      const obj = { error_name: 'Bananas' };
      const el = getElByKeyWithSpy(obj, schemaKeys);
      expect(el).toEqual('Bananas');
      expect(getElByKeyWithSpy).toHaveBeenCalledOnce();
    });    
  });

});
