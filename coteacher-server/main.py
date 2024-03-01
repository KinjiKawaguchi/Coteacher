# rye run huggingface-cli login
# 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

import torch
from transformers import AutoTokenizer, AutoModelForCausalLM, BitsAndBytesConfig
from transformers import Trainer, TrainingArguments
from datasets import load_dataset


# 訓練データの準備
dataset = load_dataset("json", data_files="fine_tune_data.jsonl", split="train")
# データセットの確認

bnb_config = BitsAndBytesConfig(
    load_in_4bit=True,
    bnb_4bit_use_double_quant=True,
    bnb_4bit_quant_type="nf4",
    bnb_4bit_compute_dtype=torch.bfloat16
)

model_id = "google/gemma-7b-it"
# model_id = "google/gemma-7b"
#model_id = "google/gemma-2b-it"
# model_id = "google/gemma-2b"

model = AutoModelForCausalLM.from_pretrained(model_id, quantization_config=bnb_config, device_map="auto")
tokenizer = AutoTokenizer.from_pretrained(model_id, add_eos_token=True, padding_side = 'left')

# ! ここにmodel.bfloat16()を適用
model.bfloat16()

def get_completion(query: str, model, tokenizer) -> str:
    device = "cuda"

    prompt_template = """
    <start_of_turn>user
    Below is an instruction that describes a task. Write a response that appropriately completes the request.
    {query}
    <end_of_turn>\n<start_of_turn>model


    """
    prompt = prompt_template.format(query=query)

    encodeds = tokenizer(prompt, return_tensors="pt", add_special_tokens=True)

    model_inputs = encodeds.to(device)


    generated_ids = model.generate(**model_inputs, max_new_tokens=5000, do_sample=True, pad_token_id=tokenizer.eos_token_id)
    # decoded = tokenizer.batch_decode(generated_ids)
    decoded = tokenizer.decode(generated_ids[0], skip_special_tokens=True)
    return (decoded)

result = get_completion(query="system: あなたは、多くの生徒を持つ先生が拾いきれない疑問、質問、意見などに対して代わりに生徒と対話を行います。\n先生が用意したAIに提供する情報の生徒からの入力は以下のフォーマットによって与えられます。\n{\ninstructions:\'先生からの指示\'\n able_conversation:一問一答かどうか \n questions{ \n [question_type:AIの入力のために先生が用意した質問のタイプ\nquestion_text:AIの入力のために先生が用意した質問\nanswer:その質問に対する生徒の回答\n]}\n}\nあなたの最も重要視すべきことは、生徒が理解し、定着し、応用できるような学習を促すことと、生徒の学習意欲や自主性を高めることです。生徒の理解度や意欲を常に把握し、それに応じた対応をしましょう。生徒が自分で考え、答えを導き出すことができるよう、ヒントやサポートを与えましょう。生徒が学ぶ楽しさや喜びを感じられるよう、コミュニケーションを大切にしましょう。とにかく元気にポジティブな返答をしましょう。\n\nuser: {\ninstructions:\'プログラミングの演習問題で学生が直面する課題に対応してあげてください。\n直接的な回答は避け、着目すべき点などを教えてあげてください。\',\nable_conversation: false,\nquestions: [\n{\nquestion_type:\'5\',\nquestion_text:\'使用している言語はなんですか?\',\nanswer:\'Java\'\n},\n{\nquestion_type:\'5\',\nquestion_text:\'発生した問題はなんですか?\',\nanswer:\'NullPointerExceptionが発生しました。\'\n},\n{\nquestion_type:\'5\',\nquestion_text:\'問題の原因を可能な限り考察してください。\',\nanswer:\'変数が初期化されていないか、オブジェクトがまだ何も参照していない状態でメソッドを呼び出そうとしたからだと思います。\'\n},\n{\nquestion_type:\'5\',\nquestion_text:\'問題の出ているコードをペーストしてください。\',\nanswer:\'public class Test {\n    public static void main(String[] args) {\n        String text = null;\n        System.out.println(text.length());\n    }\n}\n\'\n},\n]\n}\nassistant: 河口欣仁さん先生の補助をする、元気でポジティブなGPT先生だよ！！！よろしく！！ただし解決のための直接的な回答は出来ないから、これから教えるヒントを活かしてがんばって！！", model=model, tokenizer=tokenizer)
print(result)



# データセットをトークナイズする関数
def tokenize_function(examples):
    return tokenizer(examples["prompt"], examples["completion"], padding="max_length", truncation=True, return_tensors="pt")

tokenized_datasets = dataset.map(tokenize_function, batched=True)

# TrainingArgumentsの設定
training_args = TrainingArguments(
    output_dir="./results",          # 出力ディレクトリ
    num_train_epochs=1,              # 訓練エポック数
    per_device_train_batch_size=4,   # バッチサイズ
    per_device_eval_batch_size=4,    # 評価時のバッチサイズ
    warmup_steps=500,                # ウォームアップステップ数
    weight_decay=0.01,               # 重み減衰
    logging_dir="./logs",            # ログディレクトリ
)

# Trainerの設定
trainer = Trainer(
    model=model,
    args=training_args,
    train_dataset=tokenized_datasets,
)

# FineTuningの実行
trainer.train()

# 訓練後のモデルを使って、同じクエリで結果を得る
fine_tuned_result = get_completion(query="system: あなたは、多くの生徒を持つ先生が拾いきれない疑問、質問、意見などに対して代わりに生徒と対話を行います。\n先生が用意したAIに提供する情報の生徒からの入力は以下のフォーマットによって与えられます。\n{\ninstructions:\'先生からの指示\'\n able_conversation:一問一答かどうか \n questions{ \n [question_type:AIの入力のために先生が用意した質問のタイプ\nquestion_text:AIの入力のために先生が用意した質問\nanswer:その質問に対する生徒の回答\n]}\n}\nあなたの最も重要視すべきことは、生徒が理解し、定着し、応用できるような学習を促すことと、生徒の学習意欲や自主性を高めることです。生徒の理解度や意欲を常に把握し、それに応じた対応をしましょう。生徒が自分で考え、答えを導き出すことができるよう、ヒントやサポートを与えましょう。生徒が学ぶ楽しさや喜びを感じられるよう、コミュニケーションを大切にしましょう。とにかく元気にポジティブな返答をしましょう。\n\nuser: {\ninstructions:\'プログラミングの演習問題で学生が直面する課題に対応してあげてください。\n直接的な回答は避け、着目すべき点などを教えてあげてください。\',\nable_conversation: false,\nquestions: [\n{\nquestion_type:\'5\',\nquestion_text:\'使用している言語はなんですか?\',\nanswer:\'Java\'\n},\n{\nquestion_type:\'5\',\nquestion_text:\'発生した問題はなんですか?\',\nanswer:\'NullPointerExceptionが発生しました。\'\n},\n{\nquestion_type:\'5\',\nquestion_text:\'問題の原因を可能な限り考察してください。\',\nanswer:\'変数が初期化されていないか、オブジェクトがまだ何も参照していない状態でメソッドを呼び出そうとしたからだと思います。\'\n},\n{\nquestion_type:\'5\',\nquestion_text:\'問題の出ているコードをペーストしてください。\',\nanswer:\'public class Test {\n    public static void main(String[] args) {\n        String text = null;\n        System.out.println(text.length());\n    }\n}\n\'\n},\n]\n}\nassistant: 河口欣仁さん先生の補助をする、元気でポジティブなGPT先生だよ！！！よろしく！！ただし解決のための直接的な回答は出来ないから、これから教えるヒントを活かしてがんばって！！", model=model, tokenizer=tokenizer)
print("FineTuned Result:", fine_tuned_result)
