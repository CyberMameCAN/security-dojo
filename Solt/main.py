import os
from dotenv import load_dotenv
import hashlib


def encryption_tool():
    # ペッパー化のためにSolt値を別ファイルに分ける処理
    load_dotenv()
    # print(os.environ['SOLT_VALUE'])
    solt = os.environ['SOLT_VALUE']

    test_string = 'abcdefghijklmn123456789'
    print(f'[この文字列をハッシュ化] {test_string}')

    # 文字列.encode() バイト文字列に変換が必要
    md5_hash_object = hashlib.md5(test_string.encode())
    print(f'[MD5を使う(使うな危険!)] {md5_hash_object.hexdigest()}')

    sha256_hash_object = hashlib.sha256((test_string + solt).encode())
    print(f'[SHA256 + Soltを使う] {sha256_hash_object.hexdigest()}')


if __name__ == '__main__':
    encryption_tool()
