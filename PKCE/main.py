import base64
import hashlib
from dataclasses import dataclass


@dataclass
class ConfirmPkce:
    """
    PKCEの理解のために軽く確認
    (参考) 情報処理安全確保支援士 令和5年春 午後2 問2
           RFC7636
    """
    client_code_challenge = None  # 認可要求の時に送るチャレンジコード
    auth_code_challenge = None  # アクセストークン要求の時に送るチャレンジコード

    def client_side(self):
        """
        [スマホ等]
        検証コードとチャレンジコードの生成
        
        """
        code_challenge_method = 'SHA256'

        code_verifier = 'dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk'
        print(f'検証コード: {code_verifier}')

        sha256hash = hashlib.sha256(code_verifier.encode()).digest()
        # print(f'sha256 hash: {sha256hash}')

        # ここでチャレンジコードを送ったことにする
        self.client_code_challenge = base64.urlsafe_b64encode(sha256hash).rstrip(b'=').decode()
        # print(f'チャレンジコード: {self.client_code_challenge}')

        return code_verifier, code_challenge_method

    def authorization_side(self, code_verifier, code_challenge_method):
        """
        [認可サーバ]
        検証コードを受け取って、前もって受け取っていたチャレンジコードとの関係を検証する
        """
        if code_challenge_method == 'SHA256':
            sha256hash = hashlib.sha256(code_verifier.encode()).digest()
            # print(f'sha256 hash: {sha256hash}')

            self.auth_code_challenge = base64.urlsafe_b64encode(sha256hash).rstrip(b'=').decode()
            # print(f'base64url encode: {self.auth_code_challenge}')
        else:
            self.auth_code_challenge = None

        self._compare()

    def _compare(self):
        """比較"""
        if self.client_code_challenge == self.auth_code_challenge:
            print('等しい')
        else:
            print('等しくない')




if __name__ == '__main__':
    confirm_pkce = ConfirmPkce()
    
    code_verifier, code_challenge_method = confirm_pkce.client_side()
    confirm_pkce.authorization_side(code_verifier, code_challenge_method)


