import hashlib
import base64
import struct
import io
import sys

# 参照 https://qiita.com/TakahikoKawasaki/items/8f0e422c7edd2d220e06
#
# OIDC (Open ID Connect)
#   技術的には**IDトークンを発行する**ための仕様（と覚えても良さげ）
# IDトークン
#   ヘッダー.ペイロード.署名 (JSON.JSON.バイナリー)
#   いつ・どこで・何のために。改竄を検知できる。JWT形式。

class IdTokenTool:
    def base64url_decode(self, input_str):
        """文字列をbase64でデコードする"""
        # Replace URL-safe characters and add padding
        input_str = input_str.replace('-', '+').replace('_', '/')
        padding = '=' * (4 - len(input_str) % 4)
        input_str += padding
        # Decode the string
        decoded_bytes = base64.b64decode(input_str)

        return decoded_bytes


    def hexdump_binary_to_decimal(self, binary_data):
        """バイナリデータを10進数表示する"""
        for i in range(0, len(binary_data), 16):
            chunk = binary_data[i:i+16]
            decimal_values = ' '.join(f'{b:3}' for b in chunk)
            print(f'{i:08x}: {decimal_values}')


if __name__ == '__main__':
    # _hello_world = 'SGVsbG8tV29ybGQ_'
    _header = 'eyJraWQiOiIxZTlnZGs3IiwiYWxnIjoiUlMyNTYifQ'
    _payload = 'ewogImlzcyI6ICJodHRwOi8vc2VydmVyLmV4YW1wbGUuY29tIiwKICJzdWIiOiAiMjQ4Mjg5NzYxMDAxIiwKICJhdWQiOiAiczZCaGRSa3F0MyIsCiAibm9uY2UiOiAibi0wUzZfV3pBMk1qIiwKICJleHAiOiAxMzExMjgxOTcwLAogImlhdCI6IDEzMTEyODA5NzAsCiAibmFtZSI6ICJKYW5lIERvZSIsCiAiZ2l2ZW5fbmFtZSI6ICJKYW5lIiwKICJmYW1pbHlfbmFtZSI6ICJEb2UiLAogImdlbmRlciI6ICJmZW1hbGUiLAogImJpcnRoZGF0ZSI6ICIwMDAwLTEwLTMxIiwKICJlbWFpbCI6ICJqYW5lZG9lQGV4YW1wbGUuY29tIiwKICJwaWN0dXJlIjogImh0dHA6Ly9leGFtcGxlLmNvbS9qYW5lZG9lL21lLmpwZyIKfQ'
    _signature = 'rHQjEmBqn9Jre0OLykYNnspA10Qql2rvx4FsD00jwlB0Sym4NzpgvPKsDjn_wMkHxcp6CilPcoKrWHcipR2iAjzLvDNAReF97zoJqq880ZD1bwY82JDauCXELVR9O6_B0w3K-E7yM2macAAgNCUwtik6SjoSUZRcf-O5lygIyLENx882p6MtmwaL1hd6qn5RZOQ0TLrOYu0532g9Exxcm-ChymrB4xLykpDj3lUivJt63eEGGN6DH5K6o33TcxkIjNrCD4XB1CKKumZvCedgHHF3IAK4dVEDSUoGlH9z4pP_eWYNXvqQOjGs-rDaQzUHl6cQQWNiDpWOl_lxXjQEvQ'

    idtokentool = IdTokenTool()

    decoded_bytes = idtokentool.base64url_decode(_header)
    print('ヘッダー JSON:', decoded_bytes.decode('utf-8'))

    decoded_bytes = idtokentool.base64url_decode(_payload)
    print('\nペイロード JSON:', decoded_bytes.decode('utf-8'))

    # print(decoded_bytes)
    print('\n署名:')
    decoded_bytes = idtokentool.base64url_decode(_signature)
    idtokentool.hexdump_binary_to_decimal(decoded_bytes)
